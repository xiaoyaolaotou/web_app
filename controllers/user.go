package controllers

import (
	"fmt"
	"net/http"
	"web_app/logic"
	"web_app/models"
	"web_app/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数较验
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数有误,直接返回响应
		fmt.Println(err)
		// 判断err是不是validator 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": utils.RemoveTopStruct(errs.Translate(utils.Trans)), // 翻译错误为中文
		})
		return

	}

	// 2. 业务处理
	if err := logic.Singup(&p); err != nil {
		zap.L().Error("logic sigup failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}
	// 3. 返回响应

	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
