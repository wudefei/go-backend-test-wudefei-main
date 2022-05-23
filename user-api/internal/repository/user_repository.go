package repository

import (
	"errors"
	"fmt"
	"time"
	"user-api/internal/config"
	"user-api/internal/model"
	"user-api/internal/repository/dao"
	"user-api/internal/util"
)

type UserOptional struct {
	ID   int64
	Name string
}

func (op *UserOptional) buildSQL() (string, []interface{}, error) {
	var where = ""
	var value = []interface{}{}

	buildHelper := util.NewSqlBuildHelper()

	if op.ID > 0 {
		where += buildHelper.AddCondition(value, " AND ")
		where += " id = ?"
		value = append(value, op.ID)
	}
	if op.Name != "" {
		where += buildHelper.AddCondition(value, " AND ")
		where += " name = ?"
		value = append(value, op.Name)
	}
	if where == "" {
		return where, value, errors.New("Invalid")
	}

	return where, value, nil
}

type UserRepository struct {
	logID   string
	userDao *dao.UserDao
}

func NewUserRepository(logId string) *UserRepository {
	return &UserRepository{
		logID:   logId,
		userDao: dao.NewUserDao(logId),
	}
}

func (u *UserRepository) FindOne(optional UserOptional) (*model.User, error) {
	if config.Config.Env == "MOCK" {
		return &model.User{
			ID:          1,
			Name:        "evan",
			Birth:       time.Now(),
			Address:     "shenzhen",
			Description: "boy",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			IsDeleted:   0,
		}, nil
	}
	where, value, err := optional.buildSQL()
	if err != nil {
		return nil, err
	}

	result, err := u.userDao.Query(where, value...)
	if len(result) > 1 {
		return nil, errors.New("record max than one")
	}
	if len(result) == 0 {
		return nil, nil
	}
	return result[0], nil
}

func (u *UserRepository) Save(user *model.User) error {
	if config.Config.Env == "MOCK" {
		return nil
	}

	if user.ID == 0 {
		if err := u.userDao.Create(user); err != nil {
			fmt.Println(u.logID, "create record fail:", err)
			return err
		}
	} else {
		if err := u.userDao.Save(user); err != nil {
			fmt.Println(u.logID, "save record fail:", err)
			return err

		}
	}

	return nil
}
