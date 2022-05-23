package dao

import (
	"fmt"
	"user-api/internal/db"
	"user-api/internal/model"

	"github.com/jinzhu/gorm"
)

type UserDao struct {
	logID string
}

func NewUserDao(logId string) *UserDao {
	return &UserDao{
		logID: logId,
	}
}

func (u *UserDao) Query(where interface{}, value ...interface{}) ([]*model.User, error) {
	var err error
	var userList []*model.User
	err = db.GetGinBase().Where(where, value...).Find(&userList).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println(u.logID, "not exist")
			return nil, nil
		}
		fmt.Println(u.logID, "query err:", err)
		return nil, err

	}
	return userList, nil
}

func (u *UserDao) Create(user *model.User) error {
	err := db.GetGinBase().Create(user).Error
	if err != nil {
		fmt.Println(u.logID, "create err:", err)
		return err
	}
	return nil

}

func (u *UserDao) Save(user *model.User) error {
	err := db.GetGinBase().Save(user).Error
	if err != nil {
		fmt.Println(u.logID, "save err:", err)
		return err
	}
	return nil
}
