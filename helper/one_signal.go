package helper

// import (
// 	"net/http"
// 	"io/ioutil"
// 	"encoding/json"
// 	"strings"
// )

type (
	Message struct {
		AppName        string      `json:"app_name"`
		Title          string      `json:"title"`
		Message        string      `json:"message"`
		AdditionalData interface{} `json:"additional_data"`
		PlayerID       []string    `json:"player_id"`
		Segments       []string    `json:"segments"`
		Url            string      `json:"url"`
		WebUrl         string      `json:"web_url"`
		AppUrl         string      `json:"app_url"`
	}
)

// func SendNotification(message Message) (interface{}, error) {
// 	messageJson, err := json.Marshal(message)
//     if err != nil {
//         return nil, err
//     }
// 	requestBody := ioutil.NopCloser(strings.NewReader(messageJson))

// 	request := &http.Request{
// 		Method: "POST",
// 		URL:    "",
// 		Header: map[string][]string{
// 			"Content-Type":  {"application/json; charset=UTF-8"},
// 		},
// 		Body: requestBody,
// 	}

// 	resp, err := http.DefaultClient.Do(request)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer resp.Body.Close()

// 	return resp.Body, nil
// }
