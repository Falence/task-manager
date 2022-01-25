package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/falence/taskmanager/common"
	"github.com/falence/taskmanager/data"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Handler for HTTP Post - "/tasks"
// Insert a new Task document
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var dataResource TaskResource
	// Decode the incoming Task json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Task data",
			500,
		)
		return
	}
	task := &dataResource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{C: c}

	// Insert a task document
	repo.Create(task)
	if j, err := json.Marshal(TaskResource{Data: *task}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

// Handler for HTTP Get - "/tasks"
// Returns all Task documents
func GetTasks(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{C: c}
	tasks := repo.GetAll()
	j, err := json.Marshal(TaskResource{Data: tasks})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

