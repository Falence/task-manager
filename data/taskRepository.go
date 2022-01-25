package data

import (
	"time"

	"github.com/falence/taskmanager/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TaskRepository struct {
	C *mgo.Collection
}

func (r *TaskRepository) Create(task *models.Task) error {
	obj_id := bson.NewObjectId()
	task.Id = obj_id
	task.CreatedOn = time.Now()
	task.Status = "Created"
	err := r.C.Insert(&task)
	return err
}

func (r *TaskRepository) Update(task *models.Task) error {
	// Partial update on mongodb
	err := r.C.Update(bson.M{"_id": task.Id},
		bson.M{"$set": bson.M{
			"name":        task.Name,
			"description": task.Description,
			"due":         task.Due,
			"status":      task.Status,
			"tags":        task.Tags,
		}})
	return err
}



func (r *TaskRepository) GetAll() []models.Task {}

func (r *TaskRepository) GetById(id string) (task models.Task, err error) {}

func (r *TaskRepository) GetByUser(user string) []models.Task {}
