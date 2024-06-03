package controllers

import (
	"CallTheRoll/api/services"
	"CallTheRoll/pkg/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UploadExcel 上传 Excel 文件
func UploadExcel() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, resp.Fail(http.StatusBadRequest, "文件上传失败"))
			return
		}

		if err := c.SaveUploadedFile(file, "./uploads/"+file.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, resp.Error("保存文件失败"))
			return
		}

		if err := services.ProcessExcel("./uploads/" + file.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, resp.Error("处理 Excel 文件失败"))
			return
		}

		c.JSON(http.StatusOK, resp.Success("文件上传并处理成功"))
	}
}
