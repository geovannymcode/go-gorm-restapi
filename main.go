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
	s := r.PathPrefix("/api").Subrouter()

	//user
	s.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	s.HandleFunc("/users/{id}", routes.GetUserByIdHandler).Methods("GET")
	s.HandleFunc("/users", routes.CreatedUserHandler).Methods("POST")
	s.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	s.HandleFunc("/users/{id}", routes.UpdateUserHandler).Methods("PUT")

	//task
	s.HandleFunc("/tasks", routes.GetAllTasksHandler).Methods("GET")
	s.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	s.HandleFunc("/tasks/{id}", routes.GetTaskByIdHandler).Methods("GET")
	s.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	http.ListenAndServe(":4000", r)
}
