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
	"github.com/Muntaha369/Go_REST/internals/storage/sqlite"
)

func main() {
	cfg := config.Mustload() //gets the all the variables return from the config.go file

	storage, err := sqlite.New(cfg) //creates a db file at the mentioned path in cfg

	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage Initialized", slog.String("env", cfg.Env))

	router := http.NewServeMux() //its creates a router which help to identify which handler to serve

	router.HandleFunc("POST /api/createUser", rest.New(storage))      //routes
	router.HandleFunc("GET /api/getUser/{id}", rest.GetById(storage)) //routes
	router.HandleFunc("GET /api/getUsers", rest.GetByList(storage))   //routes

	server := http.Server{ //creates a server
		Addr:    cfg.Addr, //mentiones which server to listen on
		Handler: router,   //mentions routers to handle the request
	}

	fmt.Println("server started")

	done := make(chan os.Signal, 1) //Creates a Channel which will only take os signal and its buffer size is one

	signal.Notify(done, os.Interrupt, syscall.SIGTERM) //when any CTRL c or z is recived it is sent to the done channel

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start")
		}
	}()

	<-done //this syntax blocks the the until the channel has some value

	slog.Info("Shuting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //context.Background() creates an empty context with no deadline or value and it never cancels, context.withTimeout() creates a context with an deadline here it is 5 seconds
	//*to understand context more refer pkg.go.dev

	defer cancel() //it is use to clear value in case if some resource are left locked but most of the time it cancels if you are using WithTimeout() if you are using WithCancel() you need explicitly mention cancel

	err = server.Shutdown(ctx) //This is used to shutdown the server gracefully
	if err != nil {
		slog.Error("Failed to shut down the server", "error", err)
	}

	slog.Info("Shutdown server Gracefully")
}
