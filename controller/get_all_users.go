package controller

import (
	"encoding/json"
	"net/http"
	"user-management/model"
)

func (c *Controller) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	/*var users []model.User
		session, err := mongo.Connect(error)
		if err != nil {
			panic(err)
		}
		d := session.DB("user_management").C("user")
		err = d.Find("Find").All(&users)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}*/
	var users []model.User

	if err := c.session.DB("user_management").C("user").FindId("Find").all(&users); err != nil {
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(users)
}
