package controllers

import (
	"greebel.core.be/core"
	"greebel.core.be/helper"
	"greebel.core.be/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"fmt"
	"net/http"
	"strconv"
)

// BlockUser Block a user
// @Summary Block a user
// @Description Block a user
// @Tags users
// @ID users-block-block-user
// @Accept mpfd
// @Produce plain
// @Param Authorization header string true "Auth Token"
// @Param user_id path int true "user id you want to block"
// @Success 200 {object} models.UserBlock{}
// @Router /api/v1/user/block/{user_id} [post]
func BlockUser(c echo.Context) error {
	defer c.Request().Body.Close()

	userData := c.Get("user")
	token := userData.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID, _ := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))

	user := models.User{}
	if err := user.FindbyID(userID); err != nil {
		return helper.Response(http.StatusNotFound, err, "User not found")
	}

	blockedUserID, _ := strconv.Atoi(c.Param("user_id"))
	blockedUser := models.User{}
	if err := blockedUser.FindbyID(blockedUserID); err != nil {
		return helper.Response(http.StatusNotFound, err, "User not found #2")
	}

	userBlock := models.UserBlock{
		UserID:        user.ID,
		BlockedUserID: blockedUser.ID,
	}

	tx := core.App.DB.Begin()
	if err := tx.Save(&userBlock).Error; err != nil {
		tx.Rollback()
		return helper.Response(http.StatusBadRequest, err, "Error while blocking the user")
	}
	tx.Commit()

	response := helper.HttpResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data: map[string]interface{}{
			"blocked_user": blockedUser,
			"block_data":   userBlock,
		},
	}

	return c.JSON(http.StatusOK, response)
}
