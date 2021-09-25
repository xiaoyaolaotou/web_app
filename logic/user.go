package logic

import (
	"fmt"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"

	"go.uber.org/zap"
)

// 存放业务逻辑的代码

func Singup(p *models.ParamSignUp) (err error) {
	// 1. 判断用户存不存在

	if err := mysql.CheckUserExist(p.Username); err != nil {
		fmt.Println("123")
		zap.L().Error("用户已存在", zap.Error(err))
		return err
	}

	// 2. 生成UUID
	userID := snowflake.GenID()
	fmt.Println(userID)
	// 构造一个User实例
	user := &models.User{
		// UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3. 保存进数据库
	return mysql.InsertUser(user)
}

// func SingBackup(p *models.ParamBackup) (err error) {
// 	if err := mysql.Exists(p.Username); err != nil {
// 		fmt.Println("用户已存在", err)
// 		return err
// 	}

// }
