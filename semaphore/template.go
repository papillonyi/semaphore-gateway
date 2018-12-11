package semaphore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (t *Task) SetTemplate() error {
	url := "http://10.11.88.73:3000/api/project/2/templates"
	body := map[string]interface{}{
		"inventory_id":  t.inventory.Id,
		"repository_id": 6,
		"type":          "static",
		"alias":         t.uuid,
		"ssh_key_id":    5,
		"playbook":      t.Playbook,
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
