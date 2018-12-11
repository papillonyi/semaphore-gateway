package semaphore

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"net/http/cookiejar"
)

type Task struct {
	uuid         uuid.UUID
	ProjectId    int `json:"projectId"`
	RepositoryId int `json:"repositoryId"`
	InventoryId  int `json:"InventoryId"`
	Environment  environment
	Playbook     string `json:"Playbook" binding:"required"`
	Debug        bool   `json:"Debug"`
	DryRun       bool   `json:"DryRun"`
	client       *http.Client
	inventory    inventory
}

func (t *Task) Login() {
	cookieJar, _ := cookiejar.New(nil)
	url := "http://10.11.88.73:3000/api/auth/login"

	auth := map[string]string{
		"auth":     "ops",
		"password": "6FBkLc2fRN7jWH4b",
	}

	authByte, _ := json.Marshal(auth)

	req, _ := http.NewRequest("POST", url, bytes.NewReader(authByte))

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{
		Jar: cookieJar,
	}
	client.Do(req)

	t.client = client
}
