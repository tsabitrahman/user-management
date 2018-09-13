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

func (c *Controller) Signin(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var user2 model.User
	config, _ := config.GetConfig("config.yaml")

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pass := user.Password
	var sha = sha1.New()
	sha.Write([]byte(pass))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)
	user.Password = encryptedString

	db := c.MongoDB.DB(config.MongoDatabase).C("user")
	uname := user.UserName
	err := db.Find(bson.M{"username": uname}).One(&user2)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	b, _ := json.Marshal(&user2.Password)
	s := string(b)
	b1, _ := json.Marshal(&user.Password)
	s1 := string(b1)

	fmt.Println(s)
	fmt.Println(s1)
	if s == s1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(`{"message":"success"}`))
	}
}
