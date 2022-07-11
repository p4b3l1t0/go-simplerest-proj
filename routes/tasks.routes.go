package routes

import (
	"encoding/json"
	"net/http"

	"github.com/p4b3l1t0/go-simplerest-proj/db"
	"github.com/p4b3l1t0/go-simplerest-proj/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json. NewEncoder(w).Encode(tasks)

}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Task not found"))
	}

	json.NewEncoder(w).Encode(&task)



}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task []models.Task
	json.NewDecoder(r.Body).Decode(&task)
	createdTask := db.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&task)
}


func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound) // 404
		w.Write([]byte("Task not found"))
	}

	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusNoContent) // 204
}