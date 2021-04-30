package models

type User struct {
	Id          string `json:"_id" bson:"_id,omitempty"`
	UserId      string `json:"userId" bson:"userId"`
	FirstName   string `json:"firstName" bson:"firstName"`
	LastName    string `json:"lastName" bson:"lastName"`
	Email       string `json:"email" bson:"email"`
	Gender      string `json:"gender" bson:"gender"`
	UserPicture string `json:"userPicture" bson:"userPicture"`
	City        string `json:"city" bson:"city"`
	Country     string `json:"country" bson:"country"`
	DOB         string `json:"dob" bson:"dob"`
}

var UsersList []User
