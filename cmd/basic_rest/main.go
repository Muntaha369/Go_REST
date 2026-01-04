package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Muntaha369/Go_REST/internals/config"
)

func main() {
	cfg := config.Mustload()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to REST api"))
	})

	server :=http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}

	fmt.Println("server started")

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt)

	go func(){
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start")
	}
	}()

	<- done

	slog.Info("Shuting down the server")

	ctx,cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shut down the server", "error", err)
	}

	slog.Info("Shutdown server Gracefully")
}