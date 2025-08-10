package presenter

import (
	"github.com/madsilver/task-manager/internal/domain"
	"time"
)

type TaskResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	OwnerID   int64     `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToTaskResponse converts a domain.Task to a TaskResponse.
// task: the domain.Task to convert
// returns *TaskResponse
func ToTaskResponse(task *domain.Task) *TaskResponse {
	return &TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		Summary:   task.Summary,
		OwnerID:   task.OwnerID,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}

// ToTaskResponseList converts a slice of domain.Task to a slice of TaskResponse.
// tasks: the slice of domain.Task to convert
// returns []*TaskResponse
func ToTaskResponseList(tasks []*domain.Task) []*TaskResponse {
	var responses []*TaskResponse
	for _, task := range tasks {
		responses = append(responses, ToTaskResponse(task))
	}
	return responses
}
