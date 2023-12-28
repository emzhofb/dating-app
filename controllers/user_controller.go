package controllers

import (
	"dating-app/configs"
	"dating-app/middleware"
	"dating-app/models"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func RegisterUserController(c echo.Context) error {
	var userRegister models.UserCreate
	c.Bind(&userRegister)

	hashedPassword, err := HashPassword(userRegister.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Register User",
			err.Error(),
		))
	}

	var userDB models.User
	userDB.Email = userRegister.Email
	userDB.Password = string(hashedPassword)

	err = configs.DB.Create(&userDB).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, BaseResponse(
			http.StatusBadRequest,
			"Failed Register User",
			err.Error(),
		))
	}

	var userResponse = models.UserResponse{
		UserId:    userDB.UserId,
		Email:     userDB.Email,
		CreatedAt: userDB.CreatedAt,
		UpdatedAt: userDB.UpdatedAt,
	}
	return c.JSON(http.StatusCreated, BaseResponse(
		http.StatusOK,
		"Success Register User",
		userResponse,
	))
}

func LoginUserController(c echo.Context) error {
	userLogin := models.UserLogin{}
	c.Bind(&userLogin)

	var userDB models.User
	configs.DB.First(&userDB, "email", userLogin.Email)
	hashedPassword := userDB.Password

	err := VerifyPassword(hashedPassword, userLogin.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Login Customer",
			err.Error(),
		))
	}

	token, _ := middleware.GenerateJWT(userDB.UserId)

	var userResponse = models.LoginResponse{
		Email: userDB.Email,
		Token: token,
	}
	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Login Customer",
		userResponse,
	))
}

func UpdateUserController(c echo.Context) error {
	userId := middleware.GetUserIdFromJWT(c)
	var userUpdate models.Profile

	c.Bind(&userUpdate)

	var userDB models.User
	configs.DB.First(&userDB, "user_id", userId)
	userDB.Profile.Name = userUpdate.Name
	userDB.Profile.Age = userUpdate.Age
	userDB.Profile.Gender = userUpdate.Gender
	userDB.Profile.Bio = userUpdate.Bio
	userDB.Profile.Limit = 10
	userDB.Profile.IsPremium = false

	err := configs.DB.Save(&userDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Update Data",
			err.Error(),
		))
	}

	userDB.Password = ""
	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Update Data User",
		userDB,
	))
}
