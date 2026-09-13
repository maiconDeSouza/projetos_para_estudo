package router

import (
	"encoding/json"
	"net/http"
)

func StartServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		resp := map[string]string{
			"msg": "ok",
		}

		json.NewEncoder(w).Encode(resp)
	})

	return mux
}
