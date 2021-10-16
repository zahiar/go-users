package users

import (
	"go-users/src/utils"
)

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Age         int    `json:"age"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phone_number"`
}

func GetById(id int) (User, error) {
	db := utils.GetDBConnection()

	user := User{}
	err := db.First(&user, id).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
