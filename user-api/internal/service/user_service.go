package service

import (
	"errors"
	"fmt"
	"time"
	"user-api/internal/data"
	"user-api/internal/dto"
	"user-api/internal/repository"
	"user-api/internal/util"

	"user-api/internal/db/redis"
	"user-api/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go/log"
	// redigo "github.com/garyburd/redigo/redis"
)

type UserService struct {
	ctx   *gin.Context
	logID string
}

func NewUserService(ctx *gin.Context, logId string) *UserService {
	return &UserService{
		ctx:   ctx,
		logID: logId,
	}
}

/**
 * 发送手机验证码
 *
 */
func (s *UserService) SendSms(req *data.SendSmsReq) error {
	//1.生成验证码
	verifiCode := util.RandCode()
	//2.存入redis
	err := redis.NewRedisHelper(s.logID).StringWithExpire(data.VerifiCodePrefix+req.Phone, verifiCode, 1*60) //1分钟
	if err != nil {
		return err
	}
	//3.发送短信验证码
	if err = NewSmsService(s.logID).SendSms(req.Phone, verifiCode); err != nil {
		fmt.Println(s.logID, "send sms fail")
		return err
	}
	return nil
}

/**
 * 登录
 *
 */
func (s *UserService) UserLogin(req *data.UserLoginReq) (data.UserLoginRspData, error) {
	var rspData data.UserLoginRspData
	//1.校验验证码
	verifiCode, err := redis.NewRedisHelper(s.logID).StringGet(req.Phone)
	if err != nil {
		return rspData, err
	}
	//验证码过期检查----TODO
	if verifiCode == "" {
		return rspData, errors.New("Invalid verifiCode")
	}

	if verifiCode != req.VerifiCode {
		return rspData, errors.New("Invalid verifiCode")
	}

	//2.生成登录token
	token := util.RandToken()

	//3.保存token到redis
	err = redis.NewRedisHelper(s.logID).StringWithExpire(data.TokenPrefix+req.Phone, token, 5*60) //5分钟
	if err != nil {
		return rspData, err
	}

	//4.返回token
	rspData.Token = token
	return rspData, nil
}

/**
 * 查询用户信息
 *
 */
func (s *UserService) QueryUser(req *data.QueryUserReq, token string) (data.QueryUserRspData, error) {
	userRepo := repository.NewUserRepository(s.logID)
	user, err := userRepo.FindOne(repository.UserOptional{ID: req.ID, Name: req.Name})
	if err != nil {
		return data.QueryUserRspData{}, err
	}
	if user == nil {
		return data.QueryUserRspData{}, nil
	}

	//token 续期,每次续5分钟
	redis.NewRedisHelper(s.logID).StringWithExpire(data.TokenPrefix+user.Phone, token, 5*60) //5分钟

	return dto.NewDtoAssembler(s.logID).AssembleQueryUser(user)
}

/**
 * 添加用户信息
 *
 */
func (s *UserService) AddUser(req *data.AddUserReq, token string) error {
	userRepo := repository.NewUserRepository(s.logID)
	user, err := userRepo.FindOne(repository.UserOptional{Name: req.Name})
	if err != nil {
		log.Error(err)
		return err
	}

	t := time.Now()
	if user == nil {
		user = &model.User{}
		user.CreatedAt = t
	}
	user.Name = req.Name
	user.Birth, _ = time.ParseInLocation("2006-01-02", req.Birth, time.Local)
	user.Address = req.Address
	user.Description = req.Description
	user.UpdatedAt = t
	user.IsDeleted = 0

	if err = userRepo.Save(user); err != nil {
		log.Error(err)
		return err
	}

	//token 续期,每次续5分钟
	redis.NewRedisHelper(s.logID).StringWithExpire(data.TokenPrefix+user.Phone, token, 5*60) //5分钟

	return nil
}

/**
 * 更新用户信息
 *
 */
func (s *UserService) UpdateUser(req *data.UpdateUserReq, token string) error {
	userRepo := repository.NewUserRepository(s.logID)
	user, err := userRepo.FindOne(repository.UserOptional{ID: req.ID, Name: req.Name})
	if err != nil {
		log.Error(err)
		return err
	}
	if user == nil || user.IsDeleted == 1 {
		return errors.New("user not exist")
	}

	t := time.Now()
	user.Name = req.Name
	user.Birth, _ = time.ParseInLocation("2006-01-02", req.Birth, time.Local)
	user.Address = req.Address
	user.Description = req.Description
	user.UpdatedAt = t

	if err = userRepo.Save(user); err != nil {
		log.Error(err)
		return err
	}

	//token 续期,每次续5分钟
	redis.NewRedisHelper(s.logID).StringWithExpire(data.TokenPrefix+user.Phone, token, 5*60) //5分钟

	return nil
}

/**
 * 软删除用户信息
 *
 */
func (s *UserService) DeleteUser(req *data.DeleteUserReq, token string) error {
	userRepo := repository.NewUserRepository(s.logID)

	user, err := userRepo.FindOne(repository.UserOptional{ID: req.ID, Name: req.Name})
	if err != nil {
		log.Error(err)
		return err
	}
	if user == nil {
		return errors.New("user not exist")
	}
	if user.IsDeleted == 1 {
		return nil
	}

	user.UpdatedAt = time.Now()
	user.IsDeleted = 1

	if err = userRepo.Save(user); err != nil {
		log.Error(err)
		return err
	}

	//token 续期,每次续5分钟
	redis.NewRedisHelper(s.logID).StringWithExpire(data.TokenPrefix+user.Phone, token, 5*60) //5分钟

	return nil
}
