package controller

import mgo "gopkg.in/mgo.v2"

type Controller struct {
	MongoDB *mgo.Session
}
