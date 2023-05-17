package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"pet-project/internal/user"
	"pet-project/pkg/logging"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Creating router")
	router := httprouter.New()

	logger.Info("register user handler")
	handler := user.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	log.Println("start application")

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("server is listening port 0.0.0.0:1234")
	log.Fatal(server.Serve(listener))

}
