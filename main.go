package main

import (
	"net/http"

	"github.com/Geovanny0401/go-gorm-restapi/db"
	"github.com/Geovanny0401/go-gorm-restapi/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConecction()

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	http.ListenAndServe(":4000", r)
}
