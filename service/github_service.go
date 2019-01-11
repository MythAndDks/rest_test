package github_service

import (
	"net/http"
	"sample_service/api"
)

// to route services according to request
func Route_service() {
	http.HandleFunc("/create/repository", api.CreateRepo)
	http.ListenAndServe(":8080", nil)
}

