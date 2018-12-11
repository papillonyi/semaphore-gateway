package semaphore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type inventory struct {
	Id            int    `json:"id"`
	InventoryType string `json:"type"`
	Removed       bool   `json:"removed"`
	Name          string `json:"name"`
	KeyId         int    `json:"key_id"`
	SshKeyId      int    `json:"ssh_key_id"`
}

func (t *Task) SetInventory() error {
	url := "http://10.11.88.73:3000/api/project/2/inventory"
	body := map[string]interface{}{
		"inventory":  "string",
		"type":       "static",
		"name":       t.uuid,
		"ssh_key_id": 5,
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
