package knack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/buger/jsonparser"
	"injection.javamifi.com/core"
	"injection.javamifi.com/models"
)

func LoginKnack(email, password string) {

	url := "https://api.knack.com/v1/applications/5e12d832b4e5fb0015df642f/session"

	payload := strings.NewReader(`{
	  "email":"` + email + `",
	  "password":"` + password + `"
  	}`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	var token string
	var userId string
	if value, err := jsonparser.GetString([]byte(body), "session", "user", "token"); err == nil {
		token = value
	}
	if idUser, err := jsonparser.GetString([]byte(body), "session", "user", "id"); err == nil {
		userId = idUser
	}
	// update user
	type (
		UserEmailFilter struct {
			Email string `condition:"WHERE" json:"email"`
		}
	)
	tx := core.App.DB.Begin()
	checkUser := models.User{}
	type CheckUserFilter struct {
		Email string `condition:"WHERE" json:"email"`
	}
	checkUserFilter := CheckUserFilter{
		Email: email,
	}
	if checkUser.Find(&checkUserFilter); checkUser.ID == 0 {
		fmt.Println("user not found")
	}
	checkUser.UserId = userId
	checkUser.Token = token

	if err := tx.Save(&checkUser).Error; err != nil {
		tx.Rollback()
		fmt.Println(err)
		fmt.Println("failed change user")
	}
	tx.Commit()

}
