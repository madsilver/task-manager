package presenter

import "github.com/madsilver/task-manager/internal/entity"

func PopulateTask(ent *entity.Task) *Task {
	task := &Task{
		ID:      ent.ID,
		UserID:  ent.UserID,
		Summary: ent.Summary,
	}
	if ent.Date != nil {
		task.Date = *ent.Date
	}
	return task
}

func PopulateTasks(entities []*entity.Task) []*Task {
	var tasks []*Task
	for _, ent := range entities {
		tasks = append(tasks, PopulateTask(ent))
	}
	return tasks
}
