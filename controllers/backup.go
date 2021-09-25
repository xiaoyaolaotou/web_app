package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"web_app/dao/mysql"
	"web_app/models"

	"github.com/gin-gonic/gin"
)

// 判断数据是否存在

func BackupHandler(c *gin.Context) {

	var ba []*models.ParamBackup
	if err := c.ShouldBindJSON(&ba); err != nil {
		fmt.Println(err)
	}

	for _, v := range ba {
		fmt.Printf("原生的用户都有:%+v\n", v.Username)
	}

	if err := SingBacup(ba); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
		return
	}

}

// 查询
func QueryBackupHandler(c *gin.Context) {
	// var data []*models.ParamBackup
	user := QueryBackup()

	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": user,
	})

}

func SingBacup(user []*models.ParamBackup) (err error) {
	for _, v := range user {
		fmt.Println(v.Username)

		x := &models.ParamBackup{
			Username: v.Username,
			Password: v.Password,
		}

		fmt.Printf("用户都有哪些:%s\n", x.Username)

		Ex := exist(x.Username)
		if Ex {
			if err = Update(x); err != nil {
				fmt.Println("用户更新错误了哟")
				return err
			}
		} else {
			if err = Insert(x); err != nil {
				fmt.Println("插入出错")
				return err
			}
		}

	}
	return
}

// 检查数据是否存在

func exist(user string) bool {
	var username string
	err := mysql.DB.QueryRow("SELECT id FROM user WHERE username=?", user).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no data")
			return false
		} else {
			fmt.Printf("错误: %s", err)
			return false
		}
	}
	return true
}

// 插入
func Insert(user *models.ParamBackup) (err error) {
	sql := `insert into user(username,password)values(?,?) `
	// for _, v := range user {
	fmt.Println(user.Username)
	_, err = mysql.DB.Exec(sql, user.Username, user.Password)
	if err != nil {
		fmt.Println("插入失败", err)
	}
	fmt.Printf("用户：%s 已添加\n", user.Username)

	return
}

// 更新
func Update(user *models.ParamBackup) (err error) {
	sql := `update user set password=? where username=?`
	_, err = mysql.DB.Exec(sql, user.Password, user.Username)

	if err != nil {
		fmt.Println("更新失败", err)
	}
	fmt.Printf("用户：%s 已更新\n", user.Username)
	return
}

// 查询
func QueryBackup() []*models.ParamBackup {
	sql := `select username,password from user`
	var users []*models.ParamBackup
	err := mysql.DB.Select(&users, sql)
	if err != nil {
		fmt.Printf("查询所有数据失败:%+v", err)
		return nil
	}
	fmt.Println("查询出来的数据是: ", users)
	return users

}
