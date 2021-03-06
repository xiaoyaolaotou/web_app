package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"web_app/models"
)

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(username) from user where username = ?`
	var count int
	if err := DB.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	fmt.Println(count)

	return err
}

func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	password := encryptPassword(user.Password)
	fmt.Sprintf("密码是什么:%+v", password)
	// 执行SQL语句放库

	sqlStr := `insert into user(username,password)values(?,?)`
	_, err = DB.Exec(sqlStr, user.Username, password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte("asfjsjfl@LJ#"))

	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
