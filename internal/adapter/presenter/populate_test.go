package presenter

import (
	"github.com/madsilver/task-manager/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPopulateTasks(t *testing.T) {
	date := "2024-03-29 10:00:00"
	ents := []*entity.Task{
		{
			ID:      1,
			UserID:  1,
			Summary: "test",
			Date:    &date,
		},
	}

	tasks := PopulateTasks(ents)

	assert.Equal(t, len(ents), len(tasks))
	assert.Equal(t, ents[0].ID, tasks[0].ID)
	assert.Equal(t, ents[0].UserID, tasks[0].UserID)
	assert.Equal(t, ents[0].Summary, tasks[0].Summary)
	assert.Equal(t, *ents[0].Date, tasks[0].Date)
}
