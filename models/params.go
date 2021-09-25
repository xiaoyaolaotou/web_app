package models

// 定义请的参数结构体
type ParamSignUp struct {
	Username   string `json:"username,omitempty" binding:"required"`
	Password   string `json:"password,omitempty" binding:"required"`
	RePassword string `json:"re_password,omitempty" binding:"required,eqfield=Password"`
}
