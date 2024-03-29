package entity

import (
	"strconv"
)

type Task struct {
	ID      uint64
	UserID  uint64
	Summary string
	Date    *string
}

func NewTask(user any, summary string) (*Task, error) {
	userId, err := strconv.ParseUint(user.(string), 10, 64)
	if err != nil {
		return nil, err
	}

	return &Task{
		UserID:  userId,
		Summary: summary,
	}, nil
}
