package main

import (
	"fmt"
	"net/http"

	"github.com/Joshdike/mongoapi/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.GET("/user", uc.GetAllUsers)
	r.POST("/user", uc.CreateUser)
	r.PUT("/user/:id", uc.UpdateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println(err)
	}
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
