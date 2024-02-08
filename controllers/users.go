package controllers

import (
	"greebel.core.be/core"
	"greebel.core.be/helper"
	"greebel.core.be/models"

	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type (
	CreateUserRequest struct {
		Name         string `json:"name"`
		Username     string `json:"username"`
		Email        string `json:"email"`
		Password     string `json:"password"`
		Address      string `json:"address"`
		PhoneNumber  string `json:"phone_number"`
		StatusActive int    `json:"status_active"`
		IsPartner    int    `json:"is_partner"`
	}

	EditUser struct {
		Name           string `json:"name" validate:"required"`
		Username       string `json:"username" validate:"required"`
		Email          string `json:"email" validate:"required,email"`
		Company        string `json:"company"`
		Address        string `json:"address"`
		ProfilePicture string `json:"profile_picture"`
		PhoneNumber    string `json:"phone_number" validate:"required"`
	}

	EditUserRequest struct {
		User EditUser `json:"user"`
	}

	ChangePasswordRequest struct {
		Password       string `json:"password"`
		RepeatPassword string `json:"repeat_password"`
		OTP            string `json:"otp"`
	}

	RequestOTPRequest struct {
		Category string `json:"category"`
	}

	OTPFilter struct {
		Code     string `condition:"WHERE" json:"code"`
		Expired  string `condition:"WHERE" json:"expired"`
		UserID   int    `condition:"WHERE" json:"user_id"`
		Category string `condition:"WHERE" json:"category"`
	}

	CreateUserReportRequest struct {
		ReportCategoryID int `json:"report_category_id"`
	}
)

// MyProfile get my profile
// @Summary Get my profile
// @Description Get my profile
// @Tags users
// @ID users-my-profile
// @Accept json
// @Produce application/json
// @Param Authorization header string true "Auth Token"
// @Success 200 {object} helper.HttpResponse{data=models.User}
// @Router /api/v1/user/my-profile [get]
func MyProfile(c echo.Context) error {
	defer c.Request().Body.Close()

	data := c.Get("user")
	token := data.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID, _ := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))
	var complate = false
	percentage := 0
	user := models.User{}
	err := core.App.DB.Table("users").
		Where("id = ?", userID).
		First(&user).
		Error
	if err != nil {
		return helper.Response(http.StatusNotFound, err, "User not found")
	}

	if user.ProfilePicture != "" && user.ProfilePicture != "null" && user.Email != "" && user.Email != "null" && user.PhoneNumber != "" && user.PhoneNumber != "null" && user.Name != "" && user.Name != "null" && user.Username != "" && user.Username != "null" && user.Address != "" && user.Address != "null" {
		complate = true
	}

	if user.Email != "" && user.Email != "null" {
		percentage += 20
	}

	if user.PhoneNumber != "" && user.PhoneNumber != "null" {
		percentage += 20
	}

	if user.Name != "" && user.Name != "null" {
		percentage += 20
	}

	if user.Address != "" && user.Address != "null" {
		percentage += 20
	}

	if user.Username != "" && user.Username != "null" {
		percentage += 10
	}

	if user.ProfilePicture != "" && user.ProfilePicture != "null" {
		percentage += 10
	}

	var response helper.HttpResponse

	response = helper.HttpResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data: map[string]interface{}{
			"user":       user,
			"isComplate": complate,
			"percentage": percentage,
			"api_token":  user.PlayerID,
		},
	}

	return c.JSON(http.StatusOK, response)
}

// EditUserProfile Edit user profile
// @Summary Edit user profile
// @Description Edit user profile
// @Tags users
// @ID users-edit-user-profile
// @Accept json
// @Produce application/json
// @Param Authorization header string true "Auth Token"
// @Param RequestBody body EditUserRequest true "JSON Request Body"
// @Success 200 {object} models.User{}
// @Router /api/v1/user/edit-user-profile [patch]
func EditUserProfile(c echo.Context) error {
	defer c.Request().Body.Close()

	data := c.Get("user")
	token := data.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID, _ := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))

	user := models.User{}
	if err := user.FindbyID(userID); err != nil {
		return helper.Response(http.StatusNotFound, err, "User not found")
	}

	requestBody := new(EditUserRequest)
	if err := c.Bind(requestBody); err != nil {
		return helper.Response(http.StatusBadRequest, err, "Binding request error")
	}

	validate := validator.New()
	response := map[string]interface{}{}
	tx := core.App.DB.Begin()

	if err := validate.Struct(requestBody.User); err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, helper.GetValidationError(err), "Validation user error")
	}

	if user.Username != requestBody.User.Username {
		checkUser := models.User{}
		type CheckUserFilter struct {
			Username string `condition:"WHERE" json:"username"`
		}
		checkUserFilter := CheckUserFilter{
			Username: requestBody.User.Username,
		}
		if checkUser.Find(&checkUserFilter); checkUser.ID != 0 {
			return helper.Response(http.StatusUnprocessableEntity, "Validation error", "This username has already been used. Try another username.")
		}
	}

	user.Name = requestBody.User.Name
	user.Username = requestBody.User.Username
	user.PhoneNumber = requestBody.User.PhoneNumber
	user.Email = requestBody.User.Email
	user.Address = requestBody.User.Address
	user.Company = requestBody.User.Company
	user.ProfilePicture = requestBody.User.ProfilePicture

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, err, "Failed while updating user data")
	}
	response["user"] = user

	tx.Commit()

	finalResponse := helper.HttpResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    response,
	}

	return c.JSON(http.StatusOK, finalResponse)
}

// ChangePassword Change password
// @Summary Change password
// @Description Change password
// @Tags users
// @ID users-change-password
// @Accept json
// @Produce application/json
// @Param Authorization header string true "Auth Token"
// @Param RequestBody body ChangePasswordRequest true "JSON Request Body"
// @Success 200 {object} models.User{}
// @Router /api/v1/user/change-password [patch]
func ChangePassword(c echo.Context) error {
	defer c.Request().Body.Close()

	data := c.Get("user")
	token := data.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID, _ := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))

	user := models.User{}
	if err := user.FindbyID(userID); err != nil {
		return helper.Response(http.StatusNotFound, err, "User not found")
	}

	requestBody := new(ChangePasswordRequest)
	if err := c.Bind(requestBody); err != nil {
		return helper.Response(http.StatusBadRequest, err, "Binding request error")
	}

	tx := core.App.DB.Begin()

	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, helper.GetValidationError(err), "Validation change password error")
	}

	if requestBody.Password != requestBody.RepeatPassword {
		return helper.Response(http.StatusBadRequest, "password field must be equal to repeat_password field", "Error")
	}

	checkOTP := models.OTP{}
	otpFilter := OTPFilter{
		Code:     requestBody.OTP,
		Expired:  "0",
		UserID:   userID,
		Category: "change_password",
	}
	if checkOTP.Find(&otpFilter); checkOTP.ID == 0 {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, "Error", "Incorrect OTP")
	}

	checkOTP.Expired = 1
	if err := tx.Save(&checkOTP).Error; err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, err, "Error while checking OTP")
	}

	password := requestBody.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return helper.Response(http.StatusBadRequest, err, "Hash password failed")
	}
	hashString := string(hash)
	user.Password = hashString

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, err, "Error while changing new password")
	}

	tx.Commit()

	response := helper.HttpResponseData{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    user,
	}

	return c.JSON(http.StatusOK, response)
}

// RequestOTP Request OTP
// @Summary Request OTP
// @Description Request OTP
// @Tags users
// @ID users-request-otp
// @Accept json
// @Produce application/json
// @Param Authorization header string true "Auth Token"
// @Param RequestBody body RequestOTPRequest true "JSON Request Body"
// @Success 200 {object} string
// @Router /api/v1/user/request-otp [post]
func RequestOTP(c echo.Context) error {
	defer c.Request().Body.Close()

	data := c.Get("user")
	token := data.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID, _ := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))

	user := models.User{}
	if err := user.FindbyID(userID); err != nil {
		return helper.Response(http.StatusNotFound, err, "User not found")
	}

	requestBody := new(RequestOTPRequest)
	if err := c.Bind(requestBody); err != nil {
		return helper.Response(http.StatusBadRequest, err, "Binding request error")
	}

	tx := core.App.DB.Begin()

	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, helper.GetValidationError(err), "Validation change password error")
	}

	var otpCode string
	for duplicate := true; duplicate; {
		otpCode = helper.RandSeqNum(4)
		otp := models.OTP{}
		otpFilter := OTPFilter{
			Code:     otpCode,
			Expired:  "0",
			UserID:   user.ID,
			Category: requestBody.Category,
		}
		if otp.Find(&otpFilter); otp.ID != 0 {
			duplicate = true
		} else {
			duplicate = false
			otp = models.OTP{
				Code:     otpCode,
				Expired:  0,
				UserID:   user.ID,
				Category: requestBody.Category,
			}
			if err := tx.Create(&otp).Error; err != nil {
				tx.Rollback()
				return helper.Response(http.StatusBadRequest, err, "Error while creating OTP")
			}

			emailData := helper.SendGrid{
				SenderMail:    core.App.Config.SYSTEM_EMAIL,
				SenderName:    core.App.Config.SYSTEM_EMAIL_NAME,
				RecipientMail: user.Email,
				RecipientName: user.Name,
				Subject:       "OTP for ZetSend",
				HTMLContent:   fmt.Sprintf("Your OTP code is: <b>%s</b>", otp.Code),
			}

			if err := helper.SendMail(emailData); err != nil {
				tx.Rollback()
				return helper.Response(http.StatusBadRequest, err, "Error while sending OTP")
			}
		}
	}

	tx.Commit()

	response := helper.HttpResponseData{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    "Request OTP success",
	}

	return c.JSON(http.StatusOK, response)
}

// UploadProfilePhoto Upload profile photo
// @Summary Upload profile photo
// @Description Upload profile photo
// @Tags users
// @ID users-upload-profile-photo
// @Accept json
// @Produce plain
// @Param Authorization header string true "Auth Token"
// @Param image formData file true "image file"
// @Success 200 {object} string
// @Router /api/v1/user/upload-profile-photo [post]
func UploadProfilePhoto(c echo.Context) error {
	defer c.Request().Body.Close()

	data := c.Get("user")
	token := data.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID, _ := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))

	user := models.User{}
	if err := user.FindbyID(userID); err != nil {
		return helper.Response(http.StatusNotFound, err, "User not found")
	}

	dir := fmt.Sprintf("public/user-%d/profile-picture/", userID)
	objects := helper.GetObjectList(dir)
	for _, object := range objects {
		err := helper.DeleteObject(object)
		if err != nil {
			return helper.Response(http.StatusBadRequest, err, fmt.Sprintf("Error while removing old profile picture"))
		}
	}

	filenames, err := helper.UploadFile(c, dir, "image")
	if err != nil {
		return helper.Response(http.StatusBadRequest, err, "Error")
	}

	fileURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s%s", helper.BucketName, dir, filenames[0])

	tx := core.App.DB.Begin()
	user.ProfilePicture = fileURL
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, err, "Error while changing new password")
	}
	tx.Commit()

	response := helper.HttpResponse{
		Status:  http.StatusOK,
		Message: "Upload success",
		Data: map[string]interface{}{
			"filename": filenames[0],
			"url":      fileURL,
			"user":     user,
		},
	}

	return c.JSON(http.StatusOK, response)
}

// UserDelete Delete users
// @Security ApiKeyAuth
// @Summary Delete users
// @Description Delete users
// @Tags users
// @ID users-user-delete
// @Produce json
// @Param Authorization header string true "Auth Token"
// @Param id path int true "ID"
// @Success 200 {object} models.User{}
// @Router /api/v1/user/delete/{id} [delete]
func UserDelete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	if err := user.FindbyID(id); err != nil {
		return helper.Response(http.StatusUnprocessableEntity, err, "data not found")
	}

	if err := user.Delete(); err != nil {
		return helper.Response(http.StatusUnprocessableEntity, err, "failed while deleting data")

	}
	httpResponse := helper.HttpResponseData{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    user,
	}

	return c.JSON(http.StatusCreated, httpResponse)
}
