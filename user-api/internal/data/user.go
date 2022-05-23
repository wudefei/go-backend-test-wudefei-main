package data

// -----------------------------------发送短信验证码------------------------------------------------------

type SendSmsReq struct {
	Phone string `json:"phone"` //手机号码
}

// -----------------------------------登录------------------------------------------------------

type UserLoginReq struct {
	Phone      string `json:"phone"`       //手机号码
	VerifiCode string `json:"verifi_code"` //短信验证码
}

type UserLoginRspData struct {
	Token string `json:"token"`
}

// -----------------------------------query user------------------------------------------------------

type QueryUserReq struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type QueryUserRspData struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Birth       string `json:"birth"`
	Address     string `json:"address"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// --------------------------------------add user------------------------------------------------------

type AddUserReq struct {
	Name        string `json:"name"`
	Birth       string `json:"birth"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

// -------------------------------------update user-------------------------------------------------------

type UpdateUserReq struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Birth       string `json:"birth"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

// ---------------------------------------delete user-----------------------------------------------------

type DeleteUserReq struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
