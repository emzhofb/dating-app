package controllers

import (
	"dating-app/configs"
	"dating-app/middleware"
	"dating-app/models"
	"fmt"
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
			"Failed Login User",
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
		"Success Login User",
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

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Update Data User",
		userDB,
	))
}

func UpdateUserPremiumController(c echo.Context) error {
	userId := middleware.GetUserIdFromJWT(c)

	var userDB models.User
	configs.DB.First(&userDB, "user_id", userId)
	userDB.Profile.Limit = -1
	userDB.Profile.IsPremium = true

	err := configs.DB.Save(&userDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Update Data",
			err.Error(),
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Update Data User Premium",
		userDB,
	))
}

func GetAllUserController(c echo.Context) error {
	userId := middleware.GetUserIdFromJWT(c)

	var userDB []models.User
	result := configs.DB.Where("user_id != ?", userId).Find(&userDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Get Data",
			result.Error.Error(),
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Get List Users",
		userDB,
	))
}

func ResetUserLimits() {
	var userDB []models.User
	result := configs.DB.Where("is_premium != ?", false).Find(&userDB)

	if result.Error != nil {
		fmt.Println("failed get data from cron")
	}

	for _, v := range userDB {
		v.Profile.Limit = 10
		err := configs.DB.Save(v).Error

		if err != nil {
			fmt.Println("failed update user limit")
		}
	}
}
