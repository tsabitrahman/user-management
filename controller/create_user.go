package controller

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"user-management/config"
	"user-management/model"

	"gopkg.in/mgo.v2/bson"
)

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	config, _ := config.GetConfig("config.yaml")

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := c.MongoDB.DB(config.MongoDatabase).C("user")
	user.ID = bson.NewObjectId()
	pass := user.Password
	var sha = sha1.New()
	sha.Write([]byte(pass))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)
	user.Password = encryptedString
	err := db.Insert(&user)
	if err != nil {
		fmt.Printf("error %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}
