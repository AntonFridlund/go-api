package main

import (
	"context"
	"fmt"
	"main/middlewares/log"
	"main/routes"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

func main() {
	host := "localhost"
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
	server.Handler = log.LogMiddleware(server.Handler)

	// Goroutine to process log entries
	go func() {
		for logEntry := range log.LogChannel {
			os.Stdout.Write([]byte(logEntry))
		}
	}()

	// Start server in separate goroutine
	go func() {
		fmt.Println("Server started successfully")
		msg := fmt.Sprintf("Listening to: %s:%s", host, port)
		fmt.Printf("%s\n%s\n", msg, strings.Repeat("-", len(msg)))
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
		close(log.LogChannel)
	})

	// Wait for shutdown signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	fmt.Println("Received shutdown signal")

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
