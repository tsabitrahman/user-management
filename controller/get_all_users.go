package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user-management/config"
	"user-management/model"

	"gopkg.in/mgo.v2/bson"
)

func (c *Controller) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	config, _ := config.GetConfig("config.yaml")

	db := c.MongoDB.DB(config.MongoDatabase).C("user")

	err := db.Find(bson.M{}).All(&users)
	if err != nil {
		fmt.Printf("error %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(users)
}
