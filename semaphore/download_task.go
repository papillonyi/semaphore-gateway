package semaphore

import (
	"fmt"
	"github.com/google/uuid"
)

type DownloadTask struct {
	uuid            uuid.UUID
	Path            string   `json:"path"`
	DownloadAddress string   `json:"download_address"`
	Ips             []string `json:"ips"`
	Md5             string   `json:"md5"`
	Overwrite       bool     `json:"overwrite"`
	User            string   `json:"user"`
}

func (t *DownloadTask) GetNewTask() (Task, error) {
	env := environment{
		path: t.Path,
		url:  t.DownloadAddress,
		md5:  t.Md5,
		user: t.User,
	}

	fmt.Println(env)

	task := Task{
		uuid:        uuid.New(),
		ProjectId:   1,
		Environment: env,
	}

	fmt.Println(task)

	if t.Overwrite {
		task.Playbook = "file-download/download_file_force"
	} else {
		task.Playbook = "file-download/download_file"
	}

	return task, nil
}

//
//func (t *Task) createEnvironment() (environment,error) {
//	env := environment{
//
//	}
//}
