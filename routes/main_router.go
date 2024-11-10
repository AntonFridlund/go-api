package routes

import (
	"main/routes/api"
	"net/http"
)

// NewRouter creates the main router
func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	apiRouter := api.NewRouter()
	router.Handle("/api/", apiRouter)

	return router
}
