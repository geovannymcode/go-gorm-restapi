package main

import (
	"net/http"

	"github.com/Geovanny0401/go-gorm-restapi/db"
	"github.com/Geovanny0401/go-gorm-restapi/models"
	"github.com/Geovanny0401/go-gorm-restapi/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConecction()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserByIdHandler).Methods("GET")
	r.HandleFunc("/users", routes.CreatedUserHandler).Methods("POST")
	r.HandleFunc("/users", routes.DeleteUserHandler).Methods("DELETE")
	http.ListenAndServe(":4000", r)
}
