package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"greebel.core.be/core"
)

//https://documentation.onesignal.com/reference/create-notification#example-code---create-notification
func SendNotificationTest() error {
	apiKey := core.App.Config.ONE_SIGNAL_API_KEY
	appID := core.App.Config.ONE_SIGNAL_APP_ID
	requestURL, _ := url.Parse("https://onesignal.com/api/v1/notifications")
	requestBody := ioutil.NopCloser(strings.NewReader(`
		{
			"app_id": "` + appID + `",
			"contents": {"en": "hwerwerfewrfew"},
			"headings": {"en": "wqqwwqwqwq"},
			"include_external_user_ids": ["h3h3-boii"],
			"data": {
				"id": "123",
				"title": "titleeee"
			}
		}
	`))
	request := &http.Request{
		Method: "POST",
		URL:    requestURL,
		Header: map[string][]string{
			"Content-Type":  {"application/json; charset=UTF-8"},
			"Authorization": {"Basic " + apiKey},
		},
		Body: requestBody,
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func SendNotification(heading, content, notificationURL string, segments, externalUserIds []string, additionalData map[string]interface{}) error {
	for i, id := range externalUserIds {
		externalUserIds[i] = fmt.Sprintf("\"%s\"", id)
	}
	externalUserIdsString := strings.Join(strings.Fields(fmt.Sprint(externalUserIds)), ",")

	for i, segment := range segments {
		segments[i] = fmt.Sprintf("\"%s\"", segment)
	}
	segmentsString := fmt.Sprintf("[%s]", strings.Join(segments, ","))

	jsonStringData, _ := json.Marshal(additionalData)
	apiKey := core.App.Config.ONE_SIGNAL_API_KEY
	appID := core.App.Config.ONE_SIGNAL_APP_ID
	requestURL, _ := url.Parse("https://onesignal.com/api/v1/notifications")
	requestBodyString := `
		{
			"app_id": "` + appID + `",
			"contents": {"en": "` + content + `"},
			"headings": {"en": "` + heading + `"},
			"include_external_user_ids": ` + externalUserIdsString + `,
			"included_segments": ` + segmentsString + `,
			"data": ` + string(jsonStringData) + `,
			"url": "` + notificationURL + `"
		}
	`
	requestBody := ioutil.NopCloser(strings.NewReader(requestBodyString))
	request := &http.Request{
		Method: "POST",
		URL:    requestURL,
		Header: map[string][]string{
			"Content-Type":  {"application/json; charset=UTF-8"},
			"Authorization": {"Basic " + apiKey},
		},
		Body: requestBody,
	}

	resp, err := http.DefaultClient.Do(request)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	return nil
}
