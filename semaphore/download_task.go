package semaphore

type environment struct {
	path string
	url  string
	md5  string
	user string
}

//
//type template struct {
//	templateId int
//	inventory
//}

type DownloadTask struct {
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

	task := Task{
		ProjectId:   1,
		Environment: env,
	}
	return task, nil
}

//
//func (t *Task) createEnvironment() (environment,error) {
//	env := environment{
//
//	}
//}
