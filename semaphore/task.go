package semaphore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type semaphoreToken struct {
	Id      string    `json:"id"`
	Created time.Time `json:"created"`
	Expired bool      `json:"expired"`
	UserID  int       `json:"user_id"`
}

type Task struct {
	ProjectId    int `json:"projectId"`
	RepositoryId int `json:"repositoryId"`
	InventoryId  int `json:"InventoryId"`
	Environment  environment
	Playbook     string `json:"Playbook" binding:"required"`
	Debug        bool   `json:"Debug"`
	DryRun       bool   `json:"DryRun"`
	token        string
	client       *http.Client
	jar          map[string][]*http.Cookie
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

func (t *Task) GetToken() (tokens []string, err error) {
	tokens = make([]string, 0, 5)
	url := "http://10.11.88.73:3000/api/user/tokens"

	req, _ := http.NewRequest("GET", url, nil)

	res, err := t.client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var semaphoreTokens []semaphoreToken
	err = json.Unmarshal(body, &semaphoreTokens)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, element := range semaphoreTokens {
		tokens = append(tokens, element.Id)
	}

	return tokens, nil
}

func (t *Task) setTemplate() error {
	return nil
}

func (t *Task) setInventory() error {
	return nil
}
