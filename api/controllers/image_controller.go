package controllers

import (
	"CallTheRoll/api/services"
	"CallTheRoll/pkg/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GenerateImage 生成学生签到状态图片
func GenerateImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 调用生成图片的逻辑
		img, err := services.GenerateImage()
		if err != nil {
			c.JSON(http.StatusInternalServerError, resp.Error("生成图片失败"))
			return
		}
		// 不论是否生成成功，清空数据库
		if err := services.ClearDatabase(); err != nil {
			c.JSON(http.StatusInternalServerError, resp.Error("清空数据库失败"))
			return
		}
		// 成功响应
		c.JSON(http.StatusOK, resp.Success(img))
	}
}
