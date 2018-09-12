package main

import (
	"fmt"
	"log"
	"net/http"
	"user-management/config"
	"user-management/controller"
	"user-management/database/mongo"

	"github.com/gorilla/mux"
)

func main() {
	config, err := config.GetConfig("config.yaml")
	if err != nil {
		log.Fatalf("cannot read configuration file: %v", err)
	}

	mongodb, err := mongo.Connect(config.MongoUser, config.MongoPassword, config.MongoHost, config.MongoPort)
	if err != nil {
		log.Fatalf("cannot connect to mongodb: %v", err)
	}

	controller := &controller.Controller{
		MongoDB: mongodb,
	}

	addr := fmt.Sprintf(":%s", config.Port)

	r := mux.NewRouter()
	r.HandleFunc("/user", controller.GetAllUsers).Methods("GET")
	r.HandleFunc("/user/{id}", controller.GetByID).Methods("GET")
	r.HandleFunc("/user", controller.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", controller.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", controller.DeleteUser).Methods("DELETE")
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
