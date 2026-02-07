package main

import "fmt"

func double(number *int) (int, error) {
	if *number < 0 {
		return 0, fmt.Errorf("tem que ser maior que zero")
	}
	*number = *number * 2
	return *number, nil
}

func main() {
	x := 5
	result, err := double(&x)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)
	fmt.Println(x)

}
