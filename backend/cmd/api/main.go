package main

import (
	_ "backend/docs"
	"backend/internal/app"
	"context"
	"fmt"
	"log"
	"time"
)

// @title Petstore
// @version 1.0.0
// @description Petstore API implementation

// @host localhost:8888
// @BasePath /
func main() {

	a, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	go a.Serve()

	// waiting for a stop signal
	<-a.Signal()

	// create a context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// countdown to make graceful shutdown explicit
	ticker, secs := time.NewTicker(1*time.Second), 5
	for {
		select {
		// countdown
		case <-ticker.C:
			fmt.Printf("%d...\n", secs)
			secs -= 1
		// stop server gracefully
		case <-ctx.Done():
			fmt.Println("Server stopped gracefully")
			return
		}
	}
}
