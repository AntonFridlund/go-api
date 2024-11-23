package main

import (
	"context"
	"fmt"
	"main/middlewares/logger"
	"main/routes"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Server settings
	server := http.Server{
		Addr:              host + ":" + port,
		MaxHeaderBytes:    2 * 1024 * 1024,
		ReadHeaderTimeout: 3 * time.Second,
		ReadTimeout:       4 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       60 * time.Second,
		Handler:           routes.NewRouter(),
	}
	server.Handler = logger.LoggerMiddleware(server.Handler)

	// Goroutine to process log entries
	go func() {
		for logEntry := range logger.LogChannel {
			os.Stdout.Write([]byte(logEntry))
		}
	}()

	// Start server in separate goroutine
	go func() {
		fmt.Println("[INFO] Starting the server...")
		fmt.Println("[INFO] Press Ctrl+C to stop the server")
		fmt.Printf("[INFO] Listening on: http://%s:%s\n\n", host, port)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Println("Error starting server:", err)
		}
	}()

	// Handle graceful shutdown logic
	var graceWaitGroup sync.WaitGroup
	graceWaitGroup.Add(1)
	server.RegisterOnShutdown(func() {
		fmt.Println("Initiating graceful shutdown")
		defer graceWaitGroup.Done()
		close(logger.LogChannel)
	})

	// Wait for shutdown signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	fmt.Println("\nReceived shutdown signal")

	// Set timeout for graceful shutdown
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Error during shutdown:", err)
	}

	// Wait for shutdown function to complete
	graceWaitGroup.Wait()
	fmt.Println("Server gracefully stopped")
}
