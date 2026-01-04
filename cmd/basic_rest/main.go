package main

import (
	"fmt"
	"log"
	"net/http"

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
	
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start")
	}

}