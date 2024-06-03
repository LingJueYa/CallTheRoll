package controllers

import (
	"CallTheRoll/api/services"
	"CallTheRoll/pkg/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

// StudentStatusRequest 接收请求中的数据
type StudentStatusRequest struct {
	Number string `json:"number"`
	Status string `json:"status"`
}

func UpdateStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var statusReq StudentStatusRequest
		if err := c.ShouldBindJSON(&statusReq); err != nil {
			c.JSON(http.StatusBadRequest, resp.Fail(http.StatusBadRequest, "数据格式错误"))
			return
		}
		// 更新学生状态
		if err := services.UpdateStatus(statusReq.Number, statusReq.Status); err != nil {
			c.JSON(http.StatusInternalServerError, resp.Error("更新学生状态失败"))
			return
		}

		c.JSON(http.StatusOK, resp.Success("学生状态更新成功"))
	}
}

// GetStudents 获取学生列表
func GetStudents() gin.HandlerFunc {
	return func(c *gin.Context) {
		students, err := services.GetStudents()
		if err != nil {
			c.JSON(http.StatusInternalServerError, resp.Error("获取学生信息失败"))
			return
		}

		c.JSON(http.StatusOK, resp.Success(students))
	}
}
