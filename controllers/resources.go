package controllers

import (
	"github.com/falence/taskmanager/models"
)

// MODELS FOR JSON RESOURCES

// For User
type (
	// For Post - /user/register
	UserResource struct {
		Data models.User `json:"data"`
	}

	// For Post - /user/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}

	// Response for authorized user Post - /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}

	// Model for authentication
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Model for authorized user with access token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
)

// For Tasks
type (
	// For Post/Put - /tasks
	// For Get - /task/id
	TaskResource struct {
		Data models.Task `json:"data"`
	}
	// For Get - /tasks
	TasksResource struct {
		Data []models.Task `json:"data"`
	}
)

// For Notes
type (
	// For Post/Put - /notes
	NoteResource struct {
		Data NoteModel `json:"data"`
	}
	// For Get - /notes
	// For /notes/tasks/id
	NotesResource struct {
		Data []models.TaskNote `json:"data"`
	}
	// Model for a TaskNote
	NoteModel struct {
		TaskId string `json:"taskid"`
		Description string `json:"description"`
	}
)