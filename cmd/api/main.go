package main

import (
	"fmt"
	"log"
	"net/http"
	"www.github.com/Maevlava/Matatani/backend/internal/config"
	server "www.github.com/Maevlava/Matatani/backend/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Printf("Error processing config: %s\nUsing OS env", err)
	}

	mainRouter := server.NewRouter()

	mainServer := http.Server{
		Addr:    ":" + cfg.HostPort,
		Handler: mainRouter,
	}

	fmt.Println("Server listening on port ", cfg.HostPort)
	log.Fatal(mainServer.ListenAndServe())
}
