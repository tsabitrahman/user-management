package controller

import (
	"fmt"
	"net/http"
	"user-management/config"
	"user-management/database/mongo"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

//DeleteUser for delete data endpoint User
func (c *Controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	config, _ := config.GetConfig("config.yaml")

	session, _ := mongo.Connect(config.MongoUser, config.MongoPassword, config.MongoHost, config.MongoPort)
	defer session.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	db := session.DB(config.MongoDatabase).C("user")
	var selector = bson.M{"_id": bson.ObjectIdHex(id)}
	err := db.Remove(selector)
	if err != nil {
		fmt.Printf("error %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(`{"message":"success"}`))
}
