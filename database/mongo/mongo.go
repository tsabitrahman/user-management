package mongo

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

//Connect establish connection to mongodb instance and return mgo.session
func Connect(username, password, hostname, port string) (*mgo.Session, error) {
	var url string
	if username == "" && password == "" {
		url = fmt.Sprintf("mongodb://%s:%s", hostname, port)
	} else {
		url = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, hostname, port)
	}

	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	return session, nil
}
