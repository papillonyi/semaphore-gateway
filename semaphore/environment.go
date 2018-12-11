package semaphore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type environment struct {
	Id   int
	path string
	url  string
	md5  string
	user string
}

func (t *Task) SetEnvironment() error {
	url := "http://10.11.88.73:3000/api/project/2/environment"
	body := map[string]interface{}{
		"environment": map[string]interface{}{
			"path": t.Environment.path,
			"url":  t.Environment.url,
			"md5":  t.Environment.md5,
			"user": t.Environment.user,
		},
	}
	bodyByte, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(bodyByte))
	res, err := t.client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	responseBody, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	err = json.Unmarshal(responseBody, &t.inventory)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
