// Package openapi3 provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package openapi3

import (
	"time"
)

// Dates defines model for Dates.
type Dates struct {
	Due   *time.Time `json:"due"`
	Start *time.Time `json:"start"`
}

// Priority defines model for Priority.
type Priority string

// List of Priority
const (
	Priority_high   Priority = "high"
	Priority_low    Priority = "low"
	Priority_medium Priority = "medium"
	Priority_none   Priority = "none"
)

// Task defines model for Task.
type Task struct {
	Dates       *Dates    `json:"dates,omitempty"`
	Description *string   `json:"description,omitempty"`
	Id          *string   `json:"id,omitempty"`
	IsDone      *bool     `json:"is_done,omitempty"`
	Priority    *Priority `json:"priority,omitempty"`
}

// CreateTasksResponse defines model for CreateTasksResponse.
type CreateTasksResponse struct {
	Task *Task `json:"task,omitempty"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Error *string `json:"error,omitempty"`
}

// ReadTasksResponse defines model for ReadTasksResponse.
type ReadTasksResponse struct {
	Task *Task `json:"task,omitempty"`
}

// SearchTasksResponse defines model for SearchTasksResponse.
type SearchTasksResponse []Task

// CreateTasksRequest defines model for CreateTasksRequest.
type CreateTasksRequest struct {
	Dates       *Dates    `json:"dates,omitempty"`
	Description *string   `json:"description,omitempty"`
	Priority    *Priority `json:"priority,omitempty"`
}

// SearchTasksRequest defines model for SearchTasksRequest.
type SearchTasksRequest struct {
	Description *string   `json:"description"`
	IsDone      *bool     `json:"is_done"`
	Priority    *Priority `json:"priority,omitempty"`
}

// UpdateTasksRequest defines model for UpdateTasksRequest.
type UpdateTasksRequest struct {
	Dates       *Dates    `json:"dates,omitempty"`
	Description *string   `json:"description,omitempty"`
	IsDone      *bool     `json:"is_done,omitempty"`
	Priority    *Priority `json:"priority,omitempty"`
}

// SearchTaskJSONRequestBody defines body for SearchTask for application/json ContentType.
type SearchTaskJSONRequestBody SearchTasksRequest

// CreateTaskJSONRequestBody defines body for CreateTask for application/json ContentType.
type CreateTaskJSONRequestBody CreateTasksRequest

// UpdateTaskJSONRequestBody defines body for UpdateTask for application/json ContentType.
type UpdateTaskJSONRequestBody UpdateTasksRequest
