package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user-management/config"
	"user-management/model"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func (c *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	var user model.User
	config, _ := config.GetConfig("config.yaml")

	vars := mux.Vars(r)
	id := vars["id"]

	db := c.MongoDB.DB(config.MongoDatabase).C("user")
	var selector = bson.M{"_id": bson.ObjectIdHex(id)}
	err := db.Find(selector).One(&user)
	if err != nil {
		fmt.Printf("error %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}
