package knack

import (
	"fmt"
	"net/http"
	"strings"

	"injection.javamifi.com/core"
)

func UpdateUserKnack(first, last, email, entry_id string) error {

	url := "https://api.knack.com/v1/objects/object_112/records/" + entry_id

	payload := strings.NewReader(`{
	  "field_1803":{
		  "first":"` + first + `",
		  "last":"` + last + `"
		  },
		  "field_1804":"` + email + `"
  	}`)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("X-Knack-Application-Id", core.App.Config.KNACK_APP_ID)
	req.Header.Add("X-Knack-REST-API-Key", core.App.Config.KNACK_REST_API_KEY)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
