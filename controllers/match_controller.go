package controllers

import (
	"dating-app/configs"
	"dating-app/middleware"
	"dating-app/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllMatchByUserController(c echo.Context) error {
	userId := middleware.GetUserIdFromJWT(c)
	var matchesDB []models.Matches
	result := configs.DB.Where("user_id_a = ?", userId).Or("user_id_b = ?", userId).Find(&matchesDB)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			result.Error.Error(),
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get List Matches",
		matchesDB,
	))
}
