package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hytkgami/log-group-sample-go/controllers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := http.NewServeMux()
	healthCheckController := controllers.NewHealthCheckController()
	router.HandleFunc("/ping", healthCheckController.Ping)
	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
