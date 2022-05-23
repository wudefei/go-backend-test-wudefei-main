package imp

import (
	"fmt"
	"net/http"

	"user-api/internal/util"

	"user-api/internal/data"
	"user-api/internal/service"

	"github.com/gin-gonic/gin"
)

/**
 * user服务接口
 *
 */
type User struct{}

func NewUserApiImp() *User {
	return &User{}
}

/**
 * 发送短信验证码
 *
 */
func (u *User) SendSms(c *gin.Context) {
	logID := util.Uniqid()
	var err error

	defer func() {
		if p := recover(); p != nil {
			util.PrettyPanic(logID, "send sms panic:", p)
		}
	}()

	req := &data.SendSmsReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, data.ResponseBase{
			Code:    util.Invalid,
			Message: err.Error(),
		})
		return
	}

	//校验参数
	if err = NewValidator(logID).ValidatorSendSms(req); err != nil {
		c.JSON(http.StatusOK, data.ResponseBase{
			Code:    util.Invalid,
			Message: err.Error(),
		})
		return
	}

	if err = service.NewUserService(c, logID).SendSms(req); err != nil {
		fmt.Println(logID, "send sms info failed:", err.Error())
		c.JSON(http.StatusOK, data.ResponseBase{
			Code:    util.InternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data.ResponseBase{
		Code:    0,
		Message: "ok",
	})

}

/**
 * 用户登录,获取token
 *
 */
func (u *User) UserLogin(c *gin.Context) {
	logID := util.Uniqid()

	defer func() {
		if p := recover(); p != nil {
			util.PrettyPanic(logID, "login panic:", p)
		}
	}()
	req := &data.QueryUserReq{
		ID:   c.GetInt64("id"),
		Name: c.GetString("name"),
	}

	rspData, err := service.NewUserService(c, logID).UserLogin(req)
	if err != nil {
		fmt.Println(logID, "login failed:", err.Error())
		c.JSON(http.StatusOK, data.ResponseBase{
			Code:    util.InternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data.ResponseBase{
		Code:    0,
		Message: "ok",
		Data:    rspData,
	})
}

/**
 * 查询用户信息
 *
 */
func (u *User) QueryUser(c *gin.Context) {
	logID := util.Uniqid()

	defer func() {
		if p := recover(); p != nil {
			util.PrettyPanic(logID, "query user panic:", p)
		}
	}()

	req := &data.QueryUserReq{
		ID:   c.GetInt64("id"),
		Name: c.GetString("name"),
	}

	token := c.GetHeader("token")

	rspData, err := service.NewUserService(c, logID).QueryUser(req, token)
	if err != nil {
		fmt.Println(logID, "query user info failed:", err.Error())
		c.JSON(http.StatusOK, data.ResponseBase{
			Code:    util.InternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data.ResponseBase{
		Code:    0,
		Message: "ok",
		Data:    rspData,
	})
}

/**
 * 添加用户信息
 *
 */
func (u *User) AddUser(c *gin.Context) {
	logID := util.Uniqid()

	defer func() {
		if p := recover(); p != nil {
			util.PrettyPanic(logID, "add user panic:", p)
		}
	}()

	req := &data.AddUserReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, data.ResponseBase{
			Code:    util.Invalid,
			Message: err.Error(),
		})
		return
	}

	token := c.GetHeader("token")

	err := service.NewUserService(c, logID).AddUser(req, token)
	if err != nil {
		fmt.Println(logID, "add user info failed:", err.Error())
		c.JSON(http.StatusOK, data.ResponseBase{
			Code:    util.InternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data.ResponseBase{
		Code:    0,
		Message: "ok",
	})
}

/**
 * 更新用户信息
 *
 */
func (u *User) UpdateUser(c *gin.Context) {
	logID := util.Uniqid()

	defer func() {
		if p := recover(); p != nil {
			util.PrettyPanic(logID, "update user panic:", p)
		}
	}()

	req := &data.UpdateUserReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, data.ResponseBase{
			Code:    util.Invalid,
			Message: err.Error(),
		})
		return
	}

	token := c.GetHeader("token")

	err := service.NewUserService(c, logID).UpdateUser(req, token)
	if err != nil {
		fmt.Println(logID, "update user info failed:", err.Error())
		c.JSON(http.StatusOK, data.ResponseBase{
			Code:    util.InternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data.ResponseBase{
		Code:    0,
		Message: "ok",
	})
}

/**
 * 软删除用户信息
 *
 */
func (u *User) DeleteUser(c *gin.Context) {
	logID := util.Uniqid()

	defer func() {
		if p := recover(); p != nil {
			util.PrettyPanic(logID, "update user panic:", p)
		}
	}()

	req := &data.DeleteUserReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, data.ResponseBase{
			Code:    util.Invalid,
			Message: err.Error(),
		})
		return
	}

	token := c.GetHeader("token")

	err := service.NewUserService(c, logID).DeleteUser(req, token)
	if err != nil {
		fmt.Println(logID, "delete user info failed:", err.Error())
		c.JSON(http.StatusOK, data.ResponseBase{
			Code:    util.InternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data.ResponseBase{
		Code:    0,
		Message: "ok",
	})
}
