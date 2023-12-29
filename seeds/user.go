package seeds

import (
	"dating-app/configs"
	"dating-app/models"
	"strconv"
)

func CreateSeedUser() {
	for i := 0; i < 15; i++ {
		email := "john.doe" + strconv.Itoa(i) + "@gmail.com"
		if !userExists(email) {
			createUser(email)
		}
	}

	for i := 0; i < 15; i++ {
		email := "jane.doe" + strconv.Itoa(i) + "@gmail.com"
		if !userExists(email) {
			createUser(email)
		}
	}
}

func userExists(email string) bool {
	var userDB models.User
	result := configs.DB.Where("email = ?", email).First(&userDB)
	return result.RowsAffected > 0
}

func createUser(email string) {
	var userDB models.User
	userDB.Email = email
	userDB.Password = "$2a$10$24ZYrwzz6eV72aQ7J.EC2OQXFZ/Zo4a0C2A03p8xOrA6LmFAqhk2q"
	userDB.Profile.Limit = 10
	configs.DB.Create(&userDB)
}
