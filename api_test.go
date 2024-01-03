package main

import (
	"bytes"
	"encoding/json"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestNewTaskHandler tests if a new task is created.
func TestNewTaskHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/task", NewTaskHandler)
	taskId := xid.New().String()
	task := Task{
		ID:       taskId,
		Title:    "Dummy Task",
		Priority: "3",
		Status:   "Dummy Status",
		Assignee: "Dummy D",
	}
	jsonValue, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/task", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

// TestGetTasksHandler tests retrieving tasks,
func TestGetTasksHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/tasks", GetTasksHandler)
	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var tasks []Task
	err := json.Unmarshal(w.Body.Bytes(), &tasks)
	if err != nil {
		return
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, tasks)
}

// TestUpdateTaskHandler tests if a task can be updated.
func TestUpdateTaskHandler(t *testing.T) {
	r := SetUpRouter()
	r.PUT("/task/:id", UpdateTaskHandler)
	task := Task{
		ID:       `5`,
		Title:    "Documentation Update",
		Priority: "2",
		Status:   "To Do",
		Assignee: "Chris B",
	}

	jsonValue, _ := json.Marshal(task)
	reqFound, _ := http.NewRequest("PUT", "/task/"+task.ID, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("PUT", "/task/100", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestDeleteTaskHandler tests if a task can be deleted.
func TestDeleteTaskHandler(t *testing.T) {
	r := SetUpRouter()
	r.DELETE("/task/:id", DeleteTaskHandler)

	task := Task{
		ID:       `5`,
		Title:    "Documentation Update",
		Priority: "2",
		Status:   "To Do",
		Assignee: "Chris B",
	}

	jsonValue, _ := json.Marshal(task)
	reqFound, _ := http.NewRequest("DELETE", "/task/"+task.ID, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("DELETE", "/task/100", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
