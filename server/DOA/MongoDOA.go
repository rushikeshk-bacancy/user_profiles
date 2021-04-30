package DOA

import (
	"fmt"
	helper "server/helpers"
	models "server/models"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// GetAllUserFromDB - Mongo Operation For Getting All Data
func GetAllUserFromDB() error {
	res := []models.User{}
	if err := helper.Collection().Find(nil).Select(bson.M{"_id": 0}).All(&res); err != nil {
		return err
	}

	models.UsersList = res
	return nil
}

// SaveUserToDB - Mongo Operation For Inserting Data
func SaveUserToDB(user models.User) error {
	fmt.Println(user)
	uuidVar := uuid.NewV4()
	user.UserId = uuidVar.String()
	return helper.Collection().Insert(user)
}

// DeleteUserFromDB - Mongo Operation For Removing Data
func DeleteUserFromDB(id string) error {
	return helper.Collection().Remove(bson.M{"userId": id})
}

// UpdateUserInDB - Mongo Operation For Updating Data
func UpdateUserInDB(user models.User) error {
	filter := bson.M{"userId": user.UserId}
	update := bson.M{"$set": bson.M{"firstName": user.FirstName, "lastName": user.LastName, "country": user.Country, "email": user.Email, "userPicture": user.UserPicture}}
	return helper.Collection().Update(filter, update)
}

// GetSpecifiedUserDoa - Mongo Operation For Updating Data
func GetSpecifiedUserDoa(user models.User) (userOne models.User, err error) {
	res := models.User{}
	if err := helper.Collection().Find(bson.M{"userId": user.UserId}).Select(bson.M{"_id": 0}).
		One(&res); err != nil {
		return res, err
	}
	return res, nil
}
