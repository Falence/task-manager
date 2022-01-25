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



func (r *TaskRepository) GetAll() []models.Task {}

func (r *TaskRepository) GetById(id string) (task models.Task, err error) {}

func (r *TaskRepository) GetByUser(user string) []models.Task {}
