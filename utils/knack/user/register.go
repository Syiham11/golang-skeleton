package knack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"injection.javamifi.com/core"
)

func RegisterKnack(first, last, email, password string) (string, error) {

	url := "https://api.knack.com/v1/objects/object_112/records"

	payload := strings.NewReader(`{
	  "field_1803":{
		  "first":"` + first + `",
		  "last":"` + last + `"
		  },
		  "field_1804":"` + email + `",
		  "field_1805":"` + password + `"
  }
  `)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("X-Knack-Application-Id", core.App.Config.KNACK_APP_ID)
	req.Header.Add("X-Knack-REST-API-Key", core.App.Config.KNACK_REST_API_KEY)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	type UserKnack struct {
		Id string `json:"id"`
	}
	var data UserKnack
	json.Unmarshal([]byte(body), &data)

	return data.Id, nil
}
