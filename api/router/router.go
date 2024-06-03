package router

import (
	"CallTheRoll/api/controllers"
	middleware "CallTheRoll/api/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	// 文件上传
	r.POST("/upload", controllers.UploadExcel())
	// 更新签到状态
	r.POST("/status", controllers.UpdateStatus())
	// 生成图片
	r.POST("/image", controllers.GenerateImage())
	// 获取所有学生
	r.GET("/students", controllers.GetStudents())

	return r
}
