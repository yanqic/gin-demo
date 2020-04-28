package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedBy int    `json:"created_by"`
}

func AddUser(data map[string]interface{}) (*User, error) {
	user := User{
		Username:  data["username"].(string),
		Password:  data["password"].(string),
		CreatedBy: data["created_by"].(int),
	}
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsers(pageNum int, pageSize int, maps interface{}) ([]*User, error) {
	var user []*User
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return user, nil
}

func GetUserTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&User{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetUser(username string) (*User, error) {
	var user User
	err := db.Where("username = ? AND deleted_on = ? ", username, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

func CheckUserUsername(username string) (bool, error) {
	var user User
	err := db.Where("username = ? AND deleted_on = ? ", username, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func CheckUserUsernameId(username string, id int) (bool, error) {
	var user User
	err := db.Where("username = ? AND id != ? AND deleted_on = ? ", username, id, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func EditUser(id int, data map[string]interface{}) error {
	var user User

	if err := db.Where("id = ? AND deleted_on = ? ", id, 0).Find(&user).Error; err != nil {
		return err
	}
	db.Model(&user).Update(data)

	return nil
}
