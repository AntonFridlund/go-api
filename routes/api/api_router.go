package api

import (
	v1 "main/routes/api/v1"
	"net/http"
)

// NewRouter creates the API router
func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	v1Router := v1.NewRouter()
	router.Handle("/api/v1/", v1Router)

	return router
}
