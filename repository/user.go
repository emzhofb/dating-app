package repository

import (
	"dating-app/configs"
	"dating-app/models"
)

func GetUserByID(userId int) (dataresult models.UserResult, err error) {
	err = configs.DB.Model(&models.User{}).Select("*").
		Where("user_id = ?", userId).
		First(&dataresult).Error
	if err != nil {
		return dataresult, err
	}
	return
}
