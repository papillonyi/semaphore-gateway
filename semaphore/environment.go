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
	Path string `json:"path"`
	Url  string `json:"url"`
	Md5  string `json:"md5"`
	User string `json:"user"`
}

func (t *Task) SetEnvironment() error {
	url := "http://10.11.88.73:3000/api/project/2/environment"

	jsonEnv, err := json.Marshal(t.Environment)

	if err != nil {
		fmt.Println(err)
	}
	body := map[string]interface{}{
		"name": t.uuid,
		"json": string(jsonEnv),
	}
	bodyByte, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(bodyByte))
	res, err := t.client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	responseBody, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	response := new(map[string]interface{})

	fmt.Println(responseBody)

	err = json.Unmarshal(responseBody, response)

	fmt.Println(response)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
