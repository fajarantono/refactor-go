package main

import (
	"database/sql"
	"log"
	"net/http"
	"refactor-go/controllers"
	"refactor-go/driver"
	"refactor-go/utils"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDB()
	router := mux.NewRouter()

	controller := controllers.Controller{}

	router.HandleFunc("/protectedEndpoint", utils.TokenVerifyMiddleWare(controller.Protected(db))).Methods("GET")
	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/login", controller.Login(db)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
