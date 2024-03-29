package presenter

import "github.com/madsilver/task-manager/internal/entity"

func PopulateTask(ent *entity.Task) *Task {
	return &Task{
		ID:      ent.ID,
		UserID:  ent.UserID,
		Summary: ent.Summary,
		Date:    ent.Date,
	}
}

func PopulateTasks(entities []*entity.Task) []*Task {
	var tasks []*Task
	for _, ent := range entities {
		tasks = append(tasks, PopulateTask(ent))
	}
	return tasks
}
