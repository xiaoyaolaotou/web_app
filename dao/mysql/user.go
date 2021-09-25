package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"web_app/models"
)

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}

	return
}

func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	password := encryptPassword(user.Password)
	// 执行SQL语句放库

	sqlStr := `insert into user(user_id,username,password)values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte("asfjsjfl@LJ#"))

	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
