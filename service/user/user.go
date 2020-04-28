package user_service

import (
	"errors"
	"gin-demo/model"
	"gin-demo/pkg/util"
)

type User struct {
	ID         int
	Username   string
	Password   string
	CreatedOn  string
	ModifiedOn string
	CreatedBy  int
	PageNum    int
	PageSize   int
}

func (a *User) Check() (*model.User, error) {
	user, err := model.GetUser(a.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	} else if user.Password != util.EncodeMD5(a.Password) {
		return nil, errors.New("密码错误")
	}

	return user, nil
}

func (a *User) Get() (*model.User, error) {
	user, err := model.GetUser(a.Username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *User) Add() (*model.User, error) {
	item := map[string]interface{}{
		"username":   a.Username,
		"password":   util.EncodeMD5(a.Password),
		"created_by": a.CreatedBy,
	}
	username, _ := model.CheckUserUsername(a.Username)

	if username {
		return nil, errors.New("username 名字重复,请更改！")
	}

	if user, err := model.AddUser(item); err == nil {
		return user, err
	} else {
		return nil, err
	}
}

func (a *User) Edit() error {
	data := map[string]interface{}{
		"username":   a.Username,
		"password":   util.EncodeMD5(a.Password),
		"created_by": a.CreatedBy,
	}

	username, _ := model.CheckUserUsernameId(a.Username, a.ID)

	if username {
		return errors.New("username 名字重复,请更改！")
	}
	err := model.EditUser(a.ID, data)
	if err != nil {
		return err
	}

	return nil
}

func (a *User) GetAll() ([]*model.User, error) {
	if a.ID != 0 {
		maps := make(map[string]interface{})
		maps["deleted_on"] = 0
		maps["id"] = a.ID
		user, err := model.GetUsers(a.PageNum, a.PageSize, maps)
		if err != nil {
			return nil, err
		}
		return user, nil
	} else {
		user, err := model.GetUsers(a.PageNum, a.PageSize, a.getMaps())
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}

func (a *User) Count() (int, error) {
	return model.GetUserTotal(a.getMaps())
}

func (a *User) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	return maps
}
