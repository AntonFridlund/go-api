package main

import (
	"fmt"
	"main/middlewares/log"
	"main/routes"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mainRouter := routes.NewRouter()
	sizeLimit := http.MaxBytesHandler(mainRouter, 1*1024*1024)
	logger := log.LogMiddleware(sizeLimit)

	// Write log channel entries to stdout
	go func() {
		for logEntry := range log.LogChannel {
			os.Stdout.Write([]byte(logEntry))
		}
	}()

	fmt.Println("Server has started at http://localhost:" + port)
	if err := http.ListenAndServe(":"+port, logger); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
