package controllers

import (
	"dating-app/configs"
	"dating-app/middleware"
	"dating-app/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func LikeController(c echo.Context) error {
	userId := middleware.GetUserIdFromJWT(c)

	isEligible := isEligible(userId)
	if !isEligible {
		return c.JSON(http.StatusBadRequest, BaseResponse(
			http.StatusBadRequest,
			"User reach the limit",
			fmt.Errorf("user reach the limit"),
		))
	}

	like := new(models.Likes)
	like.UserIdFrom = userId
	if err := c.Bind(like); err != nil {
		return c.JSON(http.StatusBadRequest, BaseResponse(
			http.StatusBadRequest,
			"Invalid Request",
			err.Error(),
		))
	}

	if isLikeRecordExists(like.UserIdFrom, like.UserIdTo) {
		return c.JSON(http.StatusAccepted, BaseResponse(
			http.StatusAccepted,
			"User has been like the same person",
			"cannot be same",
		))
	}

	var likeDB models.Likes
	likeDB.UserIdFrom = userId
	likeDB.UserIdTo = like.UserIdTo

	err := configs.DB.Create(&likeDB).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Save Like",
			err.Error(),
		))
	}

	err = reduceLimit(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Save Match",
			err.Error(),
		))
	}

	if hasMutualLike(like.UserIdFrom, like.UserIdTo) {
		var matchesDB models.Matches
		matchesDB.UserIdA = like.UserIdFrom
		matchesDB.UserIdB = like.UserIdTo
		err := configs.DB.Create(&matchesDB).Error

		if err != nil {
			return c.JSON(http.StatusInternalServerError, BaseResponse(
				http.StatusInternalServerError,
				"Failed Save Match",
				err.Error(),
			))
		}
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Save Like",
		likeDB,
	))
}

func DisLikeController(c echo.Context) error {
	userId := middleware.GetUserIdFromJWT(c)

	isEligible := isEligible(userId)
	if !isEligible {
		return c.JSON(http.StatusBadRequest, BaseResponse(
			http.StatusBadRequest,
			"Invalid Request",
			fmt.Errorf("user reach the limit"),
		))
	}

	err := reduceLimit(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"Failed Save Match",
			err.Error(),
		))
	}

	return c.JSON(http.StatusOK, BaseResponse(
		http.StatusOK,
		"Success Save UnLike",
		"user dislike person",
	))
}

func hasMutualLike(userIDFrom, userIDTo int) bool {
	var like models.Likes
	var like2 models.Likes
	likeRepo := configs.DB

	// Check if there's a mutual like
	_ = likeRepo.Where("user_id_from = ? AND user_id_to = ?", userIDFrom, userIDTo).First(&like).Error
	_ = likeRepo.Where("user_id_from = ? AND user_id_to = ?", userIDTo, userIDFrom).First(&like2).Error

	res := like.LikeId != 0 && like2.LikeId != 0
	return res
}

func isEligible(userId int) bool {
	var userDB models.User
	configs.DB.First(&userDB, "user_id", userId)

	res := true
	if !userDB.Profile.IsPremium {
		if userDB.Profile.Limit <= 0 {
			res = false
		}
	}

	return res
}

func reduceLimit(userId int) error {
	var userDB models.User
	configs.DB.First(&userDB, "user_id", userId)

	if !userDB.Profile.IsPremium {
		userDB.Profile.Limit = userDB.Profile.Limit - 1
		err := configs.DB.Save(&userDB).Error
		return err
	}

	return nil
}

func isLikeRecordExists(userIDFrom, userIDTo int) bool {
	var like models.Likes
	_ = configs.DB.Where("user_id_from = ? AND user_id_to = ?", userIDFrom, userIDTo).First(&like)

	return like.LikeId != 0
}
