package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Muntaha369/Go_REST/internals/config"
	"github.com/Muntaha369/Go_REST/internals/http/handlers/rest"
)

func main() {
	cfg := config.Mustload()

	router := http.NewServeMux()

	router.HandleFunc("POST /api/getUser", rest.New())

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Println("server started")

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start")
		}
	}()

	<-done

	slog.Info("Shuting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shut down the server", "error", err)
	}

	slog.Info("Shutdown server Gracefully")
}
