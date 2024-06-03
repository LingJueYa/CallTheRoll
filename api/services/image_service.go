package services

import (
	"CallTheRoll/api/models"
	"bytes"
	"github.com/fogleman/gg"
	"go.uber.org/zap"
	"image/color"
	"image/png"
)

// GenerateImage 生成学生签到状态图片
func GenerateImage() ([]byte, error) {
	// 获取全部学生信息
	students, err := models.GetAllStudents()
	if err != nil {
		zap.L().Error("从数据库中获取学生信息失败", zap.Error(err))
		return nil, err
	}
	// 基于学生数量计算图片高度
	width := 800
	height := 40*len(students) + 50
	const (
		fontSize  = 24
		startX    = 40
		startY    = 40
		lineSpace = 40
	)
	// 调用 gg 库创建一个新的图形上下文
	dc := gg.NewContext(width, height)
	// 白色背景
	dc.SetRGB(1, 1, 1)
	// 清空画布
	dc.Clear()

	// 加载字体
	if err := dc.LoadFontFace("blacks.ttf", fontSize); err != nil {
		zap.L().Error("加载字体失败", zap.Error(err))
		return nil, err
	}
	// 遍历学生信息
	for i, student := range students {
		// 计算当前学生的 y 坐标
		y := startY + i*lineSpace

		// 根据签到状态设置颜色
		switch student.Status {
		case "已签到":
			dc.SetColor(color.RGBA{0, 255, 0, 255}) // Green
		case "请假":
			dc.SetColor(color.RGBA{0, 0, 0, 255}) // Black
		case "缺勤":
			dc.SetColor(color.RGBA{255, 0, 0, 255}) // Red
		default:
			dc.SetColor(color.RGBA{0, 0, 0, 255}) // Default to black
		}

		line := student.Name + " " + student.Number + " " + student.Status
		dc.DrawString(line, startX, float64(y))
	}
	// 转换为 base64 编码
	var buf bytes.Buffer
	if err := png.Encode(&buf, dc.Image()); err != nil {
		zap.L().Error("转码失败", zap.Error(err))
		return nil, err
	}

	return buf.Bytes(), nil
}

// ClearDatabase 清空数据库
func ClearDatabase() error {
	zap.L().Info("清空数据库")
	return models.ClearStudents()
}
