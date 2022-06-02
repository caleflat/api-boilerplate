package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"main.go/cmd/api/routes"
	"main.go/pkg/logger"
	"main.go/pkg/middleware"

	"github.com/joho/godotenv"
)

func main() {
	r := mux.NewRouter()

	r.Use(middleware.Log)

	routes.Handle(r)

	logrus.Printf("Server listening on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func init() {
	logger.Init()

	if err := godotenv.Load(); err != nil {
		logrus.WithError(err).Fatal("Error loading .env file")
	} else {
		logrus.Println("Successfully loaded .env file")
	}

}
