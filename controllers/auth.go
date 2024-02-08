package controllers

import (
	"greebel.core.be/core"
	"greebel.core.be/helper"
	"greebel.core.be/models"
	// "greebel.core.be/utils/rabbitmq"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/google/logger"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

type (
	User struct {
		Name            string `json:"name" validate:"required"`
		Username        string `json:"username" validate:"required"`
		Email           string `json:"email" validate:"required,email"`
		Password        string `json:"password" validate:"required"`
		ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
		Company         string `json:"company"`
		Paket           int    `json:"paket"`
		PhoneNumber     string `json:"phone_number" validate:"required"`
	}

	Partner struct {
		Name          string `json:"name" validate:"required"`
		Description   string `json:"description" validate:"required"`
		Rating        int    `json:"rating"`
		Gender        string `json:"gender"`
		CategoryID    int    `json:"category_id" validate:"required"`
		CityID        int    `json:"city_id" validate:"required"`
		BankID        int    `json:"bank_id"`
		AccountNumber string `json:"account_number"`
	}

	PartnerSocmed struct {
		SocmedID    int    `json:"socmed_id" validate:"required"`
		AccountName string `json:"account_name" validate:"required"`
		Followers   int    `json:"followers" validate:"required"`
	}

	TalentRider struct {
		Name    string `json:"name"`
		Details string `json:"details"`
	}

	Talent struct {
		StageName    string        `json:"stage_name"`
		DateOfBirth  string        `json:"date_of_birth"`
		PriceRate    int           `json:"price_rate"`
		ReligionID   int           `json:"religion_id"`
		TalentRiders []TalentRider `json:"talent_riders"`
		TalentImages []string      `json:"talent_images"`
	}

	Influencer struct {
		StageName          string              `json:"stage_name" validate:"required"`
		DateOfBirth        string              `json:"date_of_birth" validate:"required"`
		ReligionID         int                 `json:"religion_id" validate:"required"`
		MaritalID          int                 `json:"marital_id" validate:"required"`
		InfluencerServices []InfluencerService `json:"influencer_services"`
		InfluencerImages   []string            `json:"influencer_images"`
	}

	InfluencerService struct {
		ServiceID    int `json:"service_id" validate:"required"`
		TotalPost    int `json:"total_post" validate:"required"`
		PostDuration int `json:"post_duration" validate:"required"`
		PriceRate    int `json:"price_rate" validate:"required"`
	}

	Venue struct {
		PicName          string            `json:"pic_name" validate:"required"`
		PicPhoneNumber   string            `json:"pic_phone_number" validate:"required"`
		VenueVendorItems []VenueVendorItem `json:"venue_vendor_items"`
	}

	Vendor struct {
		PicName          string            `json:"pic_name" validate:"required"`
		PicPhoneNumber   string            `json:"pic_phone_number" validate:"required"`
		Open             string            `json:"open" validate:"required"`
		Close            string            `json:"close" validate:"required"`
		VenueVendorItems []VenueVendorItem `json:"venue_vendor_items"`
	}

	VenueVendorItem struct {
		Name                  string   `json:"name" validate:"required"`
		Quota                 int      `json:"quota" validate:"required"`
		Price                 int      `json:"price" validate:"required"`
		Capacity              int      `json:"capacity" validate:"required"`
		AreaSize              int      `json:"area_size" validate:"required"`
		DetailsItems          string   `json:"detail_items" validate:"required"`
		MinimumOrder          int      `json:"minimum_order" validate:"required"`
		VenueVendorItemImages []string `json:"venue_vendor_item_images"`
	}

	RegisterPartnerRequest struct {
		RegisterAs     string          `json:"register_as" validate:"required"`
		User           User            `json:"user"`
		Partner        Partner         `json:"partner"`
		PartnerSocmeds []PartnerSocmed `json:"partner_socmeds"`
		Talent         Talent          `json:"talent"`
		Influencer     Influencer      `json:"influencer"`
		Venue          Venue           `json:"venue"`
		Vendor         Vendor          `json:"vendor"`
		OTP            string          `json:"otp"`
	}

	RegisterUserRequest struct {
		User User   `json:"user"`
		OTP  string `json:"otp"`
	}

	RegisterUserRequestNoOtp struct {
		User User `json:"user"`
	}

	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
		PlayerID string `json:"player_id"`
	}

	VerifyOTPRequest struct {
		Code     string `json:"code" validate:"required"`
		UserID   int    `json:"user_id"`
		Category string `json:"category" validate:"required"`
		Email    string `json:"email"`
	}

	FacebookUserDetails struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		Token     string `json:"token"`
	}

	RequestOTPNoAuthRequest struct {
		Category string `json:"category"`
		Email    string `json:"email"`
	}

	ForgotPasswordRequest struct {
		Password       string `json:"password"`
		RepeatPassword string `json:"repeat_password"`
		OTP            string `json:"otp"`
	}
)

// VerifyOTP verify otp
// @Summary Verify OTP
// @Description Verify OTP
// @Tags auth
// @ID auth-verify-otp
// @Accept json
// @Produce application/json
// @Param RequestBody body VerifyOTPRequest true "JSON Request Body"
// @Success 200 {object} models.User{}
// @Router /api/v1/auth/verify-otp [post]
func VerifyOTP(c echo.Context) error {
	defer c.Request().Body.Close()

	requestBody := new(VerifyOTPRequest)
	if err := c.Bind(requestBody); err != nil {
		return helper.Response(http.StatusBadRequest, err, "Binding request error")
	}

	tx := core.App.DB.Begin()
	validate := validator.New()

	if err := validate.Struct(requestBody); err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, helper.GetValidationError(err), "Validation request error")
	}

	type (
		OTPFilter struct {
			Code     string `condition:"WHERE" json:"code"`
			Expired  string `condition:"WHERE" json:"expired"`
			UserID   int    `condition:"WHERE" json:"user_id"`
			Category string `condition:"WHERE" json:"category"`
			Identity string `condition:"WHERE" json:"indentity"`
			Used     string `condition:"WHERE" json:"used"`
		}
	)

	checkOTP := models.OTP{}
	otpFilter := OTPFilter{}

	if requestBody.UserID != 0 {
		otpFilter = OTPFilter{
			Code:     requestBody.Code,
			Expired:  "0",
			UserID:   requestBody.UserID,
			Category: requestBody.Category,
		}
	} else {
		otpFilter = OTPFilter{
			Code:     requestBody.Code,
			Expired:  "0",
			Category: requestBody.Category,
			Identity: requestBody.Email,
		}
	}
	if checkOTP.Find(&otpFilter); checkOTP.ID == 0 {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, "Error", "Incorrect OTP")
	}

	if requestBody.UserID != 0 {
		checkUser := models.User{}
		if err := checkUser.FindbyID(checkOTP.UserID); err != nil {
			tx.Rollback()
			return helper.Response(http.StatusBadRequest, err, "User data not found")
		}
		checkUser.StatusActive = 1
		if err := tx.Save(&checkUser).Error; err != nil {
			tx.Rollback()
			return helper.Response(http.StatusBadRequest, err, "Error while updating user status")
		}
	}

	checkOTP.Expired = 0
	checkOTP.Used = "1"
	if err := tx.Save(&checkOTP).Error; err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, err, "Error while checking OTP")
	}

	tx.Commit()

	response := helper.HttpResponseData{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    "Verify OTP Successs",
	}

	return c.JSON(http.StatusCreated, response)
}

// RegisterUser register new user account
// @Summary Register new user account
// @Description Register new user account
// @Tags auth
// @ID auth-register-user
// @Accept json
// @Produce application/json
// @Param RequestBody body RegisterUserRequest true "JSON Request Body"
// @Success 200 {object} models.User{}
// @Router /api/v1/auth/register-user [post]
func RegisterUser(c echo.Context) error {
	defer c.Request().Body.Close()

	requestBody := new(RegisterUserRequest)
	if err := c.Bind(requestBody); err != nil {
		return helper.Response(http.StatusBadRequest, err, "Binding request error")
	}

	response := map[string]interface{}{}
	tx := core.App.DB.Begin()
	validate := validator.New()

	// INSERT USER DATA START
	if err := validate.Struct(requestBody.User); err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, helper.GetValidationError(err), "Validation user error")
	}

	type (
		UserEmailFilter struct {
			Email string `condition:"WHERE" json:"email"`
		}

		UsernameFilter struct {
			Username string `condition:"WHERE" json:"username"`
		}
	)

	checkExistingUser := models.User{}
	userEmailFilter := UserEmailFilter{
		Email: requestBody.User.Email,
	}
	if checkExistingUser.Find(&userEmailFilter); checkExistingUser.ID != 0 {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, "Error", "Validation user error: email already used")
	}

	checkExistingUser = models.User{}
	usernameFilter := UsernameFilter{
		Username: requestBody.User.Username,
	}
	if checkExistingUser.Find(&usernameFilter); checkExistingUser.ID != 0 {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, "Error", "Validation user error: username already used")
	}

	type (
		OTPFilter struct {
			Code     string `condition:"WHERE" json:"code"`
			Expired  string `condition:"WHERE" json:"expired"`
			Category string `condition:"WHERE" json:"category"`
			Platform string `condition:"WHERE" json:"platform"`
			Identity string `condition:"WHERE" json:"indentity"`
			Used     string `condition:"WHERE" json:"used"`
		}
	)

	otp := models.OTP{}
	otpFilter := OTPFilter{
		Code:     requestBody.OTP,
		Expired:  "0",
		Category: "register",
		Identity: requestBody.User.Email,
		Platform: "email",
		Used:     "1",
	}
	if otp.Find(&otpFilter); otp.ID == 0 {
		return helper.Response(http.StatusBadRequest, "Error", "Wrong OTP")
	}

	password := requestBody.User.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return helper.Response(http.StatusBadRequest, err, "Hash password failed")
	}
	hashString := string(hash)

	user := models.User{
		Name:           requestBody.User.Name,
		Username:       requestBody.User.Username,
		Email:          requestBody.User.Email,
		Password:       hashString,
		Address:        "",
		PhoneNumber:    requestBody.User.PhoneNumber,
		StatusActive:   0,
		IsPartner:      0,
		ProfilePicture: "",
	}
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, err, "Error while registering user data")
	}

	response["user"] = user
	// INSERT USER DATA END

	tx.Commit()

	token, err := helper.CreateJwtToken(user.ID, user.IsPartner)
	if err != nil {
		return helper.Response(http.StatusInternalServerError, err, "Error creating token")
	}

	response["token"] = token

	return c.JSON(http.StatusCreated, response)
}

// RegisterUserNoOtp register new user account no otp
// @Summary Register new user account no otp
// @Description Register new user account no otp
// @Tags auth
// @ID auth-register-user-nootp
// @Accept json
// @Produce application/json
// @Param RequestBody body RegisterUserRequestNoOtp true "JSON Request Body"
// @Success 200 {object} models.User{}
// @Router /api/v1/auth/registeruser [post]
func RegisterUserNoOtp(c echo.Context) error {
	defer c.Request().Body.Close()

	requestBody := new(RegisterUserRequestNoOtp)
	if err := c.Bind(requestBody); err != nil {
		return helper.Response(http.StatusBadRequest, err, "Binding request error")
	}

	response := map[string]interface{}{}
	tx := core.App.DB.Begin()
	validate := validator.New()

	// INSERT USER DATA START
	if err := validate.Struct(requestBody.User); err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, helper.GetValidationError(err), "Validation user error")
	}

	type (
		UserEmailFilter struct {
			Email string `condition:"WHERE" json:"email"`
		}

		UsernameFilter struct {
			Username string `condition:"WHERE" json:"username"`
		}
	)

	checkExistingUser := models.User{}
	userEmailFilter := UserEmailFilter{
		Email: requestBody.User.Email,
	}
	if checkExistingUser.Find(&userEmailFilter); checkExistingUser.ID != 0 {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, "Error", "Validation user error: email already used")
	}

	checkExistingUser = models.User{}
	usernameFilter := UsernameFilter{
		Username: requestBody.User.Username,
	}
	if checkExistingUser.Find(&usernameFilter); checkExistingUser.ID != 0 {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, "Error", "Validation user error: username already used")
	}

	password := requestBody.User.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return helper.Response(http.StatusBadRequest, err, "Hash password failed")
	}
	hashString := string(hash)

	user := models.User{
		Name:           requestBody.User.Name,
		Username:       requestBody.User.Username,
		Email:          requestBody.User.Email,
		Password:       hashString,
		Address:        "",
		Paket:          requestBody.User.Paket,
		Company:        requestBody.User.Company,
		PhoneNumber:    requestBody.User.PhoneNumber,
		StatusActive:   0,
		IsPartner:      0,
		ProfilePicture: "",
	}
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, err, "Error while registering user data")
	}

	response["user"] = user
	// INSERT USER DATA END

	tx.Commit()

	token, err := helper.CreateJwtToken(user.ID, user.IsPartner)
	if err != nil {
		return helper.Response(http.StatusInternalServerError, err, "Error creating token")
	}

	response["token"] = token

	return c.JSON(http.StatusCreated, response)
}

// Login login account
// @Summary Login account
// @Description Login account
// @Tags auth
// @ID auth-login
// @Accept json
// @Produce application/json
// @Param RequestBody body LoginRequest true "JSON Request Body"
// @Success 200 {object} string
// @Router /api/v1/auth/login [post]
func Login(c echo.Context) error {
	defer c.Request().Body.Close()

	requestBody := new(LoginRequest)
	if err := c.Bind(requestBody); err != nil {
		return helper.Response(http.StatusBadRequest, err, "Binding request error")
	}

	type (
		UserEmailFilter struct {
			Email string `condition:"WHERE" json:"email"`
		}

		UsernameFilter struct {
			Username string `condition:"WHERE" json:"username"`
		}
	)

	user := models.User{}
	notFound := core.App.DB.Where("email = ? OR username = ?", requestBody.Username, requestBody.Username).First(&user).RecordNotFound()

	if notFound {
		return helper.Response(http.StatusUnauthorized, "Wrong username/email or password", "Wrong username/email or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password))
	if err != nil {
		return helper.Response(http.StatusUnauthorized, "Wrong username/email or password", "Wrong username/email or password")
	}

	userBanned := models.UserBanned{}
	currentDate := fmt.Sprintf("%v", time.Now().Format("2006-01-02"))
	db := core.App.DB
	notFound = db.Where("user_id = ?", user.ID).
		Where("status = 1 OR (status = 2 AND end_date >= ?)", currentDate).
		First(&userBanned).
		RecordNotFound()

	if !notFound {
		return helper.Response(http.StatusUnauthorized, "", fmt.Sprintf("%s", "Your account has been banned"))
	}

	token, err := helper.CreateJwtToken(user.ID, user.IsPartner)
	if err != nil {
		return helper.Response(http.StatusInternalServerError, err, "Error creating token")
	}

	tx := core.App.DB.Begin()
	user.PlayerID = token
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, err, "Error while updating playerID")
	}
	tx.Commit()
	// if err := rabbitmq.Publish([]byte(token), "text/plain", "session"); err != nil {
	// 	fmt.Println("publish rabbitmq failed")
	// }

	response := helper.HttpResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data: map[string]interface{}{
			"token": token,
			"user":  user,
		},
	}

	return c.JSON(http.StatusOK, response)
}

var (
	oauthConfGl = &oauth2.Config{
		ClientID:     core.App.Config.GOOGLE_CLIENT_ID,
		ClientSecret: core.App.Config.GOOGLE_CLIENT_SECRET,
		RedirectURL:  core.App.Config.GOOGLE_CLIENT_REDIRECT_URL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	oauthStateStringGl = ""
	oauthConfFb        = &oauth2.Config{
		ClientID:     core.App.Config.FACEBOOK_CLIENT_ID,
		ClientSecret: core.App.Config.FACEBOOK_CLIENT_SECRET,
		RedirectURL:  core.App.Config.FACEBOOK_CLIENT_REDIRECT_URL,
		Endpoint:     facebook.Endpoint,
		Scopes:       []string{"email"},
	}
	oauthStateStringFb = helper.RandSeq(30)
)

// GoogleLoginX google login account
// @Summary Google login account
// @Description Google login account
// @Tags auth
// @ID auth-google-login-x
// @Accept json
// @Produce application/json
// @Success 200 {object} string
// @Router /api/v1/auth/google-login-x [post]
func GoogleLoginX(c echo.Context) error {
	defer c.Request().Body.Close()

	URL, err := url.Parse(oauthConfGl.Endpoint.AuthURL)
	if err != nil {
		logger.Errorf("Parse: " + err.Error())
	}
	logger.Info(URL.String())
	parameters := url.Values{}
	parameters.Add("client_id", oauthConfGl.ClientID)
	parameters.Add("scope", strings.Join(oauthConfGl.Scopes, " "))
	parameters.Add("redirect_uri", oauthConfGl.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateStringGl)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	logger.Info(url)

	return c.JSON(http.StatusOK, url)
}

func GoogleLoginCallback(c echo.Context) error {
	defer c.Request().Body.Close()

	rand.Seed(time.Now().UnixNano())
	logger.Info("Callback-gl..")

	state := c.QueryParam("state")
	logger.Info(state)
	if state != oauthStateStringGl {
		logger.Info("invalid oauth state, expected " + oauthStateStringGl + ", got " + state + "\n")
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	code := c.QueryParam("code")
	logger.Info(code)

	if code == "" {
		logger.Errorf("Code not found..")

		// User has denied access..
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	} else {
		token, err := oauthConfGl.Exchange(oauth2.NoContext, code)
		if err != nil {
			logger.Errorf("oauthConfGl.Exchange() failed with " + err.Error() + "\n")
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}
		logger.Info("TOKEN>> AccessToken>> " + token.AccessToken)
		logger.Info("TOKEN>> Expiration Time>> " + token.Expiry.String())
		logger.Info("TOKEN>> RefreshToken>> " + token.RefreshToken)

		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
		if err != nil {
			logger.Errorf("Get: " + err.Error() + "\n")
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}
		defer resp.Body.Close()

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Errorf("ReadAll: " + err.Error() + "\n")
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		logger.Info("parseResponseBody: " + string(response) + "\n")

		type CallbackResponse struct {
			ID            string `json:"id"`
			Email         string `json:"email"`
			VerifiedEmail bool   `json:"verified_email"`
			Name          string `json:"name"`
			GivenName     string `json:"given_name"`
			FamilyName    string `json:"family_name"`
			Picture       string `json:"picture"`
			RedirectURL   string `json:"redirect_url"`
			Token         string `json:"token"`
		}

		var callbackResponse CallbackResponse
		if err := json.Unmarshal([]byte(string(response)), &callbackResponse); err != nil {
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		callbackResponse.RedirectURL = core.App.Config.GOOGLE_CLIENT_REDIRECT_URL

		type (
			UserEmailFilter struct {
				Email string `condition:"WHERE" json:"email"`
			}
		)

		checkExistingUser := models.User{}
		userEmailFilter := UserEmailFilter{
			Email: callbackResponse.Email,
		}
		if checkExistingUser.Find(&userEmailFilter); checkExistingUser.ID == 0 {
			tx := core.App.DB.Begin()

			password := helper.RandSeq(10)
			hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
			if err != nil {
				return helper.Response(http.StatusBadRequest, err, "Hash password failed")
			}
			hashString := string(hash)

			newUser := models.User{
				Name:           callbackResponse.Name,
				Email:          callbackResponse.Email,
				StatusActive:   1,
				IsPartner:      0,
				ProfilePicture: callbackResponse.Picture,
				Password:       hashString,
				Username:       fmt.Sprintf("%s-%s", helper.RandSeq(5), callbackResponse.GivenName),
			}

			if err := tx.Create(&newUser).Error; err != nil {
				tx.Rollback()
				return helper.Response(http.StatusBadRequest, err, "Error while registering user data")
			}

			tx.Commit()

			checkExistingUser = newUser
		}

		accessToken, err := helper.CreateJwtToken(checkExistingUser.ID, checkExistingUser.IsPartner)
		if err != nil {
			return helper.Response(http.StatusInternalServerError, err, "Error creating token")
		}

		callbackResponse.Token = accessToken

		return c.JSON(http.StatusOK, callbackResponse)
	}
}

// GoogleLogin google login account
// @Summary Google login account
// @Description Google login account
// @Tags auth
// @ID auth-google-login
// @Accept mpfd
// @Produce plain
// @Param access_token formData string true "Google access token"
// @Param player_id formData string false "Player ID"
// @Success 200 {object} string
// @Router /api/v1/auth/google-login [post]
func GoogleLogin(c echo.Context) error {
	defer c.Request().Body.Close()

	googleAccessToken := c.FormValue("access_token")
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(googleAccessToken))

	if err != nil {
		logger.Errorf("Get: " + err.Error() + "\n")
		return helper.Response(http.StatusBadRequest, err.Error(), "Error while trying to get user info from google")
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("ReadAll: " + err.Error() + "\n")
		return helper.Response(http.StatusBadRequest, err.Error(), "Error while trying to get user info from google #2")
	}

	logger.Info("parseResponseBody: " + string(response) + "\n")

	type CallbackResponse struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Picture       string `json:"picture"`
		RedirectURL   string `json:"redirect_url"`
		Token         string `json:"token"`
	}

	var callbackResponse CallbackResponse
	if err := json.Unmarshal([]byte(string(response)), &callbackResponse); err != nil || callbackResponse.ID == "" {
		return helper.Response(http.StatusBadRequest, err, "Error while trying to get user info from google #3")
	}

	callbackResponse.RedirectURL = core.App.Config.GOOGLE_CLIENT_REDIRECT_URL

	type (
		UserEmailFilter struct {
			Email string `condition:"WHERE" json:"email"`
		}
	)

	checkExistingUser := models.User{}
	userEmailFilter := UserEmailFilter{
		Email: callbackResponse.Email,
	}
	if checkExistingUser.Find(&userEmailFilter); checkExistingUser.ID == 0 {
		tx := core.App.DB.Begin()

		password := helper.RandSeq(10)
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			return helper.Response(http.StatusBadRequest, err, "Hash password failed")
		}
		hashString := string(hash)

		newUser := models.User{
			Name:           callbackResponse.Name,
			Email:          callbackResponse.Email,
			StatusActive:   1,
			IsPartner:      0,
			ProfilePicture: callbackResponse.Picture,
			Password:       hashString,
			Username:       fmt.Sprintf("%s-%s", helper.RandSeq(5), callbackResponse.GivenName),
		}

		if err := tx.Create(&newUser).Error; err != nil {
			tx.Rollback()
			return helper.Response(http.StatusBadRequest, err, "Error while registering user data")
		}

		tx.Commit()

		checkExistingUser = newUser
	}

	accessToken, err := helper.CreateJwtToken(checkExistingUser.ID, checkExistingUser.IsPartner)
	if err != nil {
		return helper.Response(http.StatusInternalServerError, err, "Error creating token")
	}

	tx := core.App.DB.Begin()
	checkExistingUser.PlayerID = accessToken
	if err := tx.Save(&checkExistingUser).Error; err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, err, "Error while updating player ID")
	}
	tx.Commit()

	callbackResponse.Token = accessToken

	httpResponse := helper.HttpResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data: map[string]interface{}{
			"token": accessToken,
			"user":  checkExistingUser,
		},
	}

	return c.JSON(http.StatusOK, httpResponse)
}

// FacebookLoginX facebook login account
// @Summary Facebook login account
// @Description Facebook login account
// @Tags auth
// @ID auth-facebook-login-x
// @Accept json
// @Produce application/json
// @Success 200 {object} string
// @Router /api/v1/auth/facebook-login-x [post]
func FacebookLoginX(c echo.Context) error {
	defer c.Request().Body.Close()

	url := oauthConfFb.AuthCodeURL(oauthStateStringFb)

	return c.JSON(http.StatusOK, url)
}

func FacebookLoginCallback(c echo.Context) error {
	defer c.Request().Body.Close()

	state := c.QueryParam("state")
	code := c.QueryParam("code")

	if state != oauthStateStringFb {
		return c.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
	}

	token, err := oauthConfFb.Exchange(oauth2.NoContext, code)

	if err != nil || token == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/?invalidlogin=true")
	}

	fbUserDetails, fbUserDetailsError := GetUserInfoFromFacebook(token.AccessToken)

	if fbUserDetailsError != nil {
		return helper.Response(http.StatusBadRequest, fbUserDetailsError, "Error while trying to get user info from facebook")
	}

	type (
		UserEmailFilter struct {
			Email string `condition:"WHERE" json:"email"`
		}
	)

	checkExistingUser := models.User{}
	userEmailFilter := UserEmailFilter{
		Email: fbUserDetails.Email,
	}
	if checkExistingUser.Find(&userEmailFilter); checkExistingUser.ID == 0 {
		tx := core.App.DB.Begin()

		password := helper.RandSeq(10)
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			return helper.Response(http.StatusBadRequest, err, "Hash password failed")
		}
		hashString := string(hash)

		newUser := models.User{
			Name:           fbUserDetails.Name,
			Email:          fbUserDetails.Email,
			StatusActive:   1,
			IsPartner:      0,
			ProfilePicture: fmt.Sprintf("http://graph.facebook.com/%s/picture", fbUserDetails.ID),
			Password:       hashString,
			Username:       fmt.Sprintf("%s-%s", helper.RandSeq(5), fbUserDetails.FirstName),
		}

		if err := tx.Create(&newUser).Error; err != nil {
			tx.Rollback()
			return helper.Response(http.StatusBadRequest, err, "Error while registering user data")
		}

		tx.Commit()

		checkExistingUser = newUser
	}

	accessToken, err := helper.CreateJwtToken(checkExistingUser.ID, checkExistingUser.IsPartner)
	if err != nil {
		return helper.Response(http.StatusInternalServerError, err, "Error creating token")
	}

	fbUserDetails.Token = accessToken

	return c.JSON(http.StatusOK, fbUserDetails)
}

// FacebookLogin facebook login account
// @Summary Facebook login account
// @Description Facebook login account
// @Tags auth
// @ID auth-facebook-login
// @Accept mpfd
// @Produce plain
// @Param access_token formData string true "Facebook access token"
// @Param player_id formData string false "Player ID"
// @Success 200 {object} string
// @Router /api/v1/auth/facebook-login [post]
func FacebookLogin(c echo.Context) error {
	defer c.Request().Body.Close()

	facebookAccessToken := c.FormValue("access_token")
	fbUserDetails, fbUserDetailsError := GetUserInfoFromFacebook(facebookAccessToken)

	if fbUserDetailsError != nil {
		return helper.Response(http.StatusBadRequest, fbUserDetailsError, "Error while trying to get user info from facebook")
	}

	if fbUserDetails.ID == "" {
		return helper.Response(http.StatusUnauthorized, "Unauthorized", "You can not login with this access token")
	}

	type (
		UserEmailFilter struct {
			Email string `condition:"WHERE" json:"email"`
		}
	)

	checkExistingUser := models.User{}
	userEmailFilter := UserEmailFilter{
		Email: fbUserDetails.Email,
	}
	if checkExistingUser.Find(&userEmailFilter); checkExistingUser.ID == 0 {
		tx := core.App.DB.Begin()

		password := helper.RandSeq(10)
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			return helper.Response(http.StatusBadRequest, err, "Hash password failed")
		}
		hashString := string(hash)

		newUser := models.User{
			Name:           fbUserDetails.Name,
			Email:          fbUserDetails.Email,
			StatusActive:   1,
			IsPartner:      0,
			ProfilePicture: fmt.Sprintf("http://graph.facebook.com/%s/picture", fbUserDetails.ID),
			Password:       hashString,
			Username:       fmt.Sprintf("%s-%s", helper.RandSeq(5), fbUserDetails.FirstName),
		}

		if err := tx.Create(&newUser).Error; err != nil {
			tx.Rollback()
			return helper.Response(http.StatusBadRequest, err, "Error while registering user data")
		}

		tx.Commit()

		checkExistingUser = newUser
	}

	accessToken, err := helper.CreateJwtToken(checkExistingUser.ID, checkExistingUser.IsPartner)
	if err != nil {
		return helper.Response(http.StatusInternalServerError, err, "Error creating token")
	}

	tx := core.App.DB.Begin()
	checkExistingUser.PlayerID = accessToken
	if err := tx.Save(&checkExistingUser).Error; err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, err, "Error while updating player ID")
	}
	tx.Commit()

	response := helper.HttpResponse{
		Status:  http.StatusOK,
		Message: "Login success",
		Data: map[string]interface{}{
			"user":  checkExistingUser,
			"token": accessToken,
		},
	}

	return c.JSON(http.StatusOK, response)
}

func GetUserInfoFromFacebook(token string) (FacebookUserDetails, error) {
	var fbUserDetails FacebookUserDetails
	facebookUserDetailsRequest, _ := http.NewRequest("GET", "https://graph.facebook.com/me?fields=id,name,email,picture,first_name&access_token="+token, nil)
	facebookUserDetailsResponse, facebookUserDetailsResponseError := http.DefaultClient.Do(facebookUserDetailsRequest)

	if facebookUserDetailsResponseError != nil {
		return FacebookUserDetails{}, facebookUserDetailsResponseError
	}

	decoder := json.NewDecoder(facebookUserDetailsResponse.Body)
	decoderErr := decoder.Decode(&fbUserDetails)
	defer facebookUserDetailsResponse.Body.Close()

	if decoderErr != nil {
		return FacebookUserDetails{}, decoderErr
	}

	return fbUserDetails, nil
}

// UploadImage upload image
// @Summary Upload image
// @Description Upload image
// @Tags auth
// @ID auth-upload-image
// @Accept json
// @Produce plain
// @Param image formData file true "image file"
// @Success 200 {object} string
// @Router /api/v1/auth/upload-image [post]
func UploadImage(c echo.Context) error {
	defer c.Request().Body.Close()

	filenames, err := helper.UploadFile(c, helper.TemporaryDirectory, "image")
	if err != nil {
		return helper.Response(http.StatusBadRequest, err, "Error")
	}

	fileURL := fmt.Sprintf("https://storage.googleapis.com/%s/public/temp/%s", helper.BucketName, filenames[0])

	response := helper.HttpResponse{
		Status:  http.StatusOK,
		Message: "Upload success",
		Data: map[string]interface{}{
			"filename": filenames[0],
			"url":      fileURL,
		},
	}

	return c.JSON(http.StatusOK, response)
}

// DeleteImage delete image
// @Summary Delete image
// @Description Delete image
// @Tags auth
// @ID auth-delete-image
// @Accept mpfd
// @Produce plain
// @Param filename formData string true "image filename"
// @Success 200 {object} helper.HttpResponse{}
// @Router /api/v1/auth/delete-image [delete]
func DeleteImage(c echo.Context) error {
	defer c.Request().Body.Close()

	filename := c.FormValue("filename")
	err := helper.DeleteObject(helper.TemporaryDirectory + filename)
	if err != nil {
		return helper.Response(http.StatusBadRequest, err, fmt.Sprintf("Error while deleting %s.", filename))
	}

	response := helper.HttpResponse{
		Status:  http.StatusOK,
		Message: "Success",
	}

	return c.JSON(http.StatusOK, response)
}

// RequestOTPNoAuth Request OTP
// @Summary Request OTP
// @Description Request OTP
// @Tags auth
// @ID auth-request-otp
// @Accept json
// @Produce application/json
// @Param RequestBody body RequestOTPNoAuthRequest true "JSON Request Body"
// @Success 200 {object} string
// @Router /api/v1/auth/request-otp [post]
func RequestOTPNoAuth(c echo.Context) error {
	defer c.Request().Body.Close()

	requestBody := new(RequestOTPNoAuthRequest)
	if err := c.Bind(requestBody); err != nil {
		return helper.Response(http.StatusBadRequest, err, "Binding request error")
	}

	tx := core.App.DB.Begin()

	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, helper.GetValidationError(err), "Validation change password error")
	}

	type UserFilter struct {
		Email string `condition:"WHERE" json:"email"`
	}
	userFilter := UserFilter{
		Email: requestBody.Email,
	}
	user := models.User{}
	if user.Find(&userFilter); user.ID == 0 && requestBody.Category != "register" {
		return helper.Response(http.StatusBadRequest, "Error", "This email is not registered")
	}

	name := ""
	email := ""
	if user.ID != 0 {
		name = user.Name
		email = user.Email
	} else {
		name = "User"
		email = requestBody.Email
	}
	jktTz, _ := time.LoadLocation("Asia/Jakarta")
	currentTime := time.Now().In(jktTz)
	dateNow := currentTime.Format("2006-01-02")
	otpDataLast := models.OTP{}
	otpData := []models.OTP{}
	if user.ID != 0 {
		core.App.DB.Table("otp").
			Where("user_id = ?", user.ID).
			Where("category = ?", requestBody.Category).
			Where("created_at > ? ", dateNow+" 00:00:00").
			Last(&otpDataLast)

		core.App.DB.Table("otp").
			Where("user_id = ?", user.ID).
			Where("category = ?", requestBody.Category).
			Where("created_at > ? ", dateNow+" 00:00:00").
			Find(&otpData)
	} else {
		core.App.DB.Table("otp").
			Where("indentity = ?", requestBody.Email).
			Where("category = ?", requestBody.Category).
			Where("created_at > ? ", dateNow+" 00:00:00").
			Last(&otpDataLast)

		core.App.DB.Table("otp").
			Where("indentity = ?", requestBody.Email).
			Where("category = ?", requestBody.Category).
			Where("created_at > ? ", dateNow+" 00:00:00").
			Find(&otpData)
	}
	// validasi otp 1 menit
	// add date + 1 menit
	dateData := otpDataLast.CreatedAt
	var minute = +1
	if len(otpData) == 0 {
		minute = +1
	} else {
		minute = +otpDataLast.TimerOtp
	}
	addMinute := dateData.Add(time.Duration(minute) * time.Minute)

	infoOtpsend := ""
	timer_otp := minute
	var otpCode string

	if len(otpData) == 0 {
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
					TimerOtp: 1,
					Category: requestBody.Category,
					Platform: "email",
					Identity: requestBody.Email,
				}
				if err := tx.Create(&otp).Error; err != nil {
					tx.Rollback()
					return helper.Response(http.StatusBadRequest, err, "Error while creating OTP")
				}

				emailData := helper.SendGrid{
					SenderMail:    core.App.Config.SYSTEM_EMAIL,
					SenderName:    core.App.Config.SYSTEM_EMAIL_NAME,
					RecipientMail: email,
					RecipientName: name,
					Subject:       "OTP for ZetSend",
					HTMLContent:   fmt.Sprintf("Your OTP code is: <b>%s</b>", otp.Code),
				}

				if err := helper.SendMail(emailData); err != nil {
					tx.Rollback()
					return helper.Response(http.StatusBadRequest, err, "Error while sending OTP")
				}
			}
		}
		infoOtpsend = "otp send with time 1 minute"
	} else if len(otpData) == 1 && addMinute.Unix() > currentTime.Unix() {
		infoOtpsend = "wait resend otp 1 minute"
	} else if len(otpData) == 1 && addMinute.Unix() < currentTime.Unix() || addMinute.Unix() == currentTime.Unix() {
		infoOtpsend = "otp send with time 2 minute"
		if err := tx.Model(&models.OTP{}).Updates(models.OTP{Expired: 1}).Error; err != nil {
			tx.Rollback()
			return c.JSON(http.StatusUnprocessableEntity, err)
		}
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
					TimerOtp: 2,
					Category: requestBody.Category,
					Platform: "email",
					Identity: requestBody.Email,
				}
				if err := tx.Create(&otp).Error; err != nil {
					tx.Rollback()
					return helper.Response(http.StatusBadRequest, err, "Error while creating OTP")
				}

				emailData := helper.SendGrid{
					SenderMail:    core.App.Config.SYSTEM_EMAIL,
					SenderName:    core.App.Config.SYSTEM_EMAIL_NAME,
					RecipientMail: email,
					RecipientName: name,
					Subject:       "OTP for ZetSend",
					HTMLContent:   fmt.Sprintf("Your OTP code is: <b>%s</b>", otp.Code),
				}

				if err := helper.SendMail(emailData); err != nil {
					tx.Rollback()
					return helper.Response(http.StatusBadRequest, err, "Error while sending OTP")
				}
			}
		}
		timer_otp = 2
	} else if len(otpData) == 2 && addMinute.Unix() > currentTime.Unix() {
		infoOtpsend = "wait resend otp 2 minute"
	} else if len(otpData) == 2 && addMinute.Unix() < currentTime.Unix() || addMinute.Unix() == currentTime.Unix() {
		infoOtpsend = "otp send with time 4 minute"
		if err := tx.Model(&models.OTP{}).Updates(models.OTP{Expired: 1}).Error; err != nil {
			tx.Rollback()
			return c.JSON(http.StatusUnprocessableEntity, err)
		}
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
					TimerOtp: 4,
					Category: requestBody.Category,
					Platform: "email",
					Identity: requestBody.Email,
				}
				if err := tx.Create(&otp).Error; err != nil {
					tx.Rollback()
					return helper.Response(http.StatusBadRequest, err, "Error while creating OTP")
				}

				emailData := helper.SendGrid{
					SenderMail:    core.App.Config.SYSTEM_EMAIL,
					SenderName:    core.App.Config.SYSTEM_EMAIL_NAME,
					RecipientMail: email,
					RecipientName: name,
					Subject:       "OTP for ZetSend",
					HTMLContent:   fmt.Sprintf("Your OTP code is: <b>%s</b>", otp.Code),
				}

				if err := helper.SendMail(emailData); err != nil {
					tx.Rollback()
					return helper.Response(http.StatusBadRequest, err, "Error while sending OTP")
				}
			}
		}
		timer_otp = 4

	} else if len(otpData) == 3 && addMinute.Unix() > currentTime.Unix() {
		infoOtpsend = "wait resend otp 4 minute"
	} else if len(otpData) == 3 && addMinute.Unix() < currentTime.Unix() || addMinute.Unix() == currentTime.Unix() {
		infoOtpsend = "otp send with time 8 minute"
		if err := tx.Model(&models.OTP{}).Updates(models.OTP{Expired: 1}).Error; err != nil {
			tx.Rollback()
			return c.JSON(http.StatusUnprocessableEntity, err)
		}
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
					TimerOtp: 8,
					Category: requestBody.Category,
					Platform: "email",
					Identity: requestBody.Email,
				}
				if err := tx.Create(&otp).Error; err != nil {
					tx.Rollback()
					return helper.Response(http.StatusBadRequest, err, "Error while creating OTP")
				}

				emailData := helper.SendGrid{
					SenderMail:    core.App.Config.SYSTEM_EMAIL,
					SenderName:    core.App.Config.SYSTEM_EMAIL_NAME,
					RecipientMail: email,
					RecipientName: name,
					Subject:       "OTP for ZetSend",
					HTMLContent:   fmt.Sprintf("Your OTP code is: <b>%s</b>", otp.Code),
				}

				if err := helper.SendMail(emailData); err != nil {
					tx.Rollback()
					return helper.Response(http.StatusBadRequest, err, "Error while sending OTP")
				}
			}
		}
		timer_otp = 8
	} else if len(otpData) == 4 && addMinute.Unix() > currentTime.Unix() {
		infoOtpsend = "wait resend otp 8 minute"
	} else if len(otpData) == 4 && addMinute.Unix() < currentTime.Unix() || addMinute.Unix() == currentTime.Unix() {
		infoOtpsend = "otp send with time 16 minute"
		if err := tx.Model(&models.OTP{}).Updates(models.OTP{Expired: 1}).Error; err != nil {
			tx.Rollback()
			return c.JSON(http.StatusUnprocessableEntity, err)
		}
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
				if requestBody.Category != "register" {
					otp = models.OTP{
						Code:     otpCode,
						Expired:  0,
						UserID:   user.ID,
						TimerOtp: 16,
						Category: requestBody.Category,
						Platform: "email",
						Identity: requestBody.Email,
					}
					if err := tx.Create(&otp).Error; err != nil {
						tx.Rollback()
						return helper.Response(http.StatusBadRequest, err, "Error while creating OTP")
					}
				} else {

					otp = models.OTP{
						Code:     otpCode,
						Expired:  0,
						UserID:   user.ID,
						TimerOtp: 16,
						Category: requestBody.Category,
						Platform: "email",
						Identity: requestBody.Email,
					}
					if err := tx.Create(&otp).Error; err != nil {
						tx.Rollback()
						return helper.Response(http.StatusBadRequest, err, "Error while creating OTP")
					}

				}

				emailData := helper.SendGrid{
					SenderMail:    core.App.Config.SYSTEM_EMAIL,
					SenderName:    core.App.Config.SYSTEM_EMAIL_NAME,
					RecipientMail: email,
					RecipientName: name,
					Subject:       "OTP for ZetSend",
					HTMLContent:   fmt.Sprintf("Your OTP code is: <b>%s</b>", otp.Code),
				}

				if err := helper.SendMail(emailData); err != nil {
					tx.Rollback()
					return helper.Response(http.StatusBadRequest, err, "Error while sending OTP")
				}
			}
		}
		timer_otp = 16
	} else if len(otpData) == 5 && addMinute.Unix() > currentTime.Unix() {
		infoOtpsend = "wait resend otp 16 minute"
		if err := tx.Model(&models.OTP{}).Updates(models.OTP{Expired: 1}).Error; err != nil {
			tx.Rollback()
			return c.JSON(http.StatusUnprocessableEntity, err)
		}
	} else if len(otpData) == 5 && addMinute.Unix() < currentTime.Unix() || addMinute.Unix() == currentTime.Unix() {
		infoOtpsend = "sorry you have exceeded the resend otp limit"
	} else if len(otpData) > 5 {
		infoOtpsend = "sorry you have exceeded the resend otp limit"
	}

	tx.Commit()

	response := map[string]interface{}{}
	response["status"] = http.StatusOK
	response["timer_otp"] = timer_otp
	response["message"] = infoOtpsend

	return c.JSON(http.StatusOK, response)
}

// ForgotPassword Forgot password
// @Summary Forgot password
// @Description Forgot password
// @Tags auth
// @ID auth-forgot-password
// @Accept json
// @Produce application/json
// @Param RequestBody body ForgotPasswordRequest true "JSON Request Body"
// @Success 200 {object} string
// @Router /api/v1/auth/forgot-password [post]
func ForgotPassword(c echo.Context) error {
	defer c.Request().Body.Close()

	requestBody := new(ForgotPasswordRequest)
	if err := c.Bind(requestBody); err != nil {
		return helper.Response(http.StatusBadRequest, err, "Binding request error")
	}

	tx := core.App.DB.Begin()

	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, helper.GetValidationError(err), "Validation change password error")
	}

	type (
		OTPFilter struct {
			Code     string `condition:"WHERE" json:"code"`
			Expired  string `condition:"WHERE" json:"expired"`
			Category string `condition:"WHERE" json:"category"`
			Used     string `condition:"WHERE" json:"used"`
		}
	)

	otp := models.OTP{}
	otpFilter := OTPFilter{
		Code:     requestBody.OTP,
		Expired:  "0",
		Category: "forgot_password",
		Used:     "1",
	}
	if otp.Find(&otpFilter); otp.ID == 0 {
		return helper.Response(http.StatusBadRequest, "Error", "Wrong OTP")
	}

	otp.Expired = 1
	if err := tx.Save(&otp).Error; err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, err, "Error while checking OTP")
	}

	user := models.User{}
	if err := user.FindbyID(otp.UserID); err != nil {
		return helper.Response(http.StatusBadRequest, "Error", "User data not found")
	}

	if requestBody.Password != requestBody.RepeatPassword {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, "password field must be equal to repeat_password field", "Error")
	}

	password := requestBody.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		tx.Rollback()
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

// AppleLogin Apple login account
// @Summary Apple login account
// @Description Apple login account
// @Tags auth
// @ID auth-apple-login
// @Accept mpfd
// @Produce plain
// @Param access_token formData string true "The token you got from apple login response"
// @Param player_id formData string false "Player ID"
// @Success 200 {object} string
// @Router /api/v1/auth/apple-login [post]
func AppleLogin(c echo.Context) error {
	defer c.Request().Body.Close()

	accessToken := c.FormValue("access_token")
	token, _, err := new(jwt.Parser).ParseUnverified(accessToken, jwt.MapClaims{})
	if err != nil {
		return helper.Response(http.StatusBadRequest, err, "Parse token error")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return helper.Response(http.StatusBadRequest, err, "Parse claims error")
	}

	if claims["iss"] != "https://appleid.apple.com" || claims["aud"] != "com.eventori.app" {
		return helper.Response(http.StatusBadRequest, err, "Invalid token")
	}

	type (
		UserEmailFilter struct {
			Email string `condition:"WHERE" json:"email"`
		}
	)

	checkExistingUser := models.User{}
	email := fmt.Sprintf("%s", claims["email"])
	userEmailFilter := UserEmailFilter{
		Email: email,
	}
	if checkExistingUser.Find(&userEmailFilter); checkExistingUser.ID == 0 {
		tx := core.App.DB.Begin()

		password := helper.RandSeq(10)
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			return helper.Response(http.StatusBadRequest, err, "Hash password failed")
		}
		hashString := string(hash)

		newUser := models.User{
			Name:         fmt.Sprintf("%s-%s", "Eventori", helper.RandSeq(5)),
			Email:        email,
			StatusActive: 1,
			IsPartner:    0,
			Password:     hashString,
			Username:     fmt.Sprintf("%s-%s", helper.RandSeq(5), "eventori"),
		}

		if err := tx.Create(&newUser).Error; err != nil {
			tx.Rollback()
			return helper.Response(http.StatusBadRequest, err, "Error while registering user data")
		}

		tx.Commit()

		checkExistingUser = newUser
	}

	newAccessToken, err := helper.CreateJwtToken(checkExistingUser.ID, checkExistingUser.IsPartner)
	if err != nil {
		return helper.Response(http.StatusInternalServerError, err, "Error creating token")
	}

	tx := core.App.DB.Begin()
	checkExistingUser.PlayerID = newAccessToken
	if err := tx.Save(&checkExistingUser).Error; err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, err, "Error while updating player ID")
	}
	tx.Commit()

	response := helper.HttpResponse{
		Status:  http.StatusOK,
		Message: "Login success",
		Data: map[string]interface{}{
			"user":  checkExistingUser,
			"token": newAccessToken,
		},
	}

	return c.JSON(http.StatusOK, response)
}
