package main

import (
	"fmt"
	"log"
	"user-management/config"
	"user-management/database/mongo"
)

func main() {
	config, err := config.GetConfig("config.yaml")
	if err != nil {
		log.Fatalf("cannot read configuration file: %v", err)
	}

	mongodb, err := mongo.Connect(config.MongoUser,
		config.MongoPassword, config.MongoHost, config.MongoPort)
	if err != nil {
		log.Fatalf("cannot connect to mongodb: %v", err)
	}

	buildInfo, err := mongodb.BuildInfo()
	if err != nil {
		log.Fatalf("cannot get build info: %v", err)
	}
	fmt.Println(buildInfo.Version)

}
