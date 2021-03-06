package main 

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/p4b3l1t0/go-simplerest-proj/routes"	
	"github.com/p4b3l1t0/go-simplerest-proj/models"	
	"github.com/p4b3l1t0/go-simplerest-proj/db"	
)


func main() {
	
	db.DBconnect()
	
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
	
	r := mux.NewRouter()

	//endpoints
	r.HandleFunc("/", routes.HomeHandler)

	// User routes
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")

	// Task Routes
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))

}