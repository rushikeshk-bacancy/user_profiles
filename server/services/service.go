package services

import (
	// "fmt"

	Doa "server/DOA"
	models "server/models"
)

func NewUser() models.User {
	return models.User{}
}

func GetUsersList() []models.User {
	// if models.UsersList == nil || len(models.UsersList) == 0 {
	// 	// models.UsersList = []models.User{}
	// 	// return models.UsersList
	// 	Doa.GetAllUserFromDB()
	// }
	Doa.GetAllUserFromDB()
	return models.UsersList
}

// AddNewUser New User Service
func AddNewUser(user models.User) {
	// Mongo DB
	Doa.SaveUserToDB(user)
	Doa.GetAllUserFromDB()
}

// UpdateUserService Update User Service
func UpdateUserService(user models.User) {
	// MongoDb
	Doa.UpdateUserInDB(user)
	Doa.GetAllUserFromDB()

}

// RemoveUserService Remove User Service
func RemoveUserService(user models.User) {
	Doa.DeleteUserFromDB(user.UserId)
	Doa.GetAllUserFromDB()
}

// GetSpecifiedUser() Get One User Service
func GetSpecifiedUser(userOne models.User) (user models.User) {
	data, err := Doa.GetSpecifiedUserDoa(userOne)
	if err != nil {
		return models.User{}
	} else {
		return data
	}
}
