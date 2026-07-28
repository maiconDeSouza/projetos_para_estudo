package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Status string

const (
	ONLINE  Status = "Online"
	OFFLINE Status = "Offline"
	ERROR   Status = "Erro"
)

type Result struct {
	URL     string
	Status  Status
	Latency time.Duration
	Err     error
}

func (r Result) String() string {
	if r.Err != nil {
		return fmt.Sprintf("[❌ %-7s] %-30s | Latência: %-10s | Erro: %v", r.Status, r.URL, "N/A", r.Err)
	}
	return fmt.Sprintf("[✅ %-7s] %-30s | Latência: %-10s", r.Status, r.URL, r.Latency)
}

func worker(ctx context.Context, jobs <-chan string, results chan<- Result) {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	for url := range jobs {
		start := time.Now()

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			results <- Result{
				URL:     url,
				Status:  ERROR,
				Latency: time.Since(start),
				Err:     err,
			}
			continue
		}

		resp, err := client.Do(req)
		latency := time.Since(start)
		if err != nil {
			results <- Result{
				URL:     url,
				Status:  ERROR,
				Latency: latency,
				Err:     err,
			}
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			results <- Result{
				URL:     url,
				Status:  ONLINE,
				Latency: latency,
				Err:     nil,
			}
			continue
		}

		results <- Result{
			URL:     url,
			Status:  OFFLINE,
			Latency: latency,
			Err:     nil,
		}
	}
}

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://httpbin.org/delay/5",
		"https://site-que-nao-existe-123.com",
		"https://httpbin.org/status/404",
		"https://httpbin.org/status/200",
		"https://httpbin.org/delay/2",
		"https://httpbin.org/delay/4",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jobs := make(chan string, len(urls))
	results := make(chan Result, len(urls))

	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Go(func() {
			worker(ctx, jobs, results)
		})
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result.String())
	}
}
