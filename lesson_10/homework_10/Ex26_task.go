package main

type Task struct {
	title       string
	description string
	status      string
}

type Project struct {
	Tasks []Task
}

func (p *Project) AddTask(title, description, status string) {
	task := Task{
		title:       title,
		description: description,
		status:      status,
	}

	p.Tasks = append(p.Tasks, task)
}

func (p Project) CountTasksByStatus(status string) int {
	var taskAmount int

	for _, v := range p.Tasks {
		if status == v.status {
			taskAmount++
		}
	}

	return taskAmount
}
