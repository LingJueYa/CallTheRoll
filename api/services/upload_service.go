package services

import (
	"CallTheRoll/api/models"
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize"
	"go.uber.org/zap"
)

// ProcessExcel 处理 Excel 文件
func ProcessExcel(filePath string) error {
	zap.L().Info("打开 Excel 文件", zap.String("filePath", filePath))
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		zap.L().Error("打开 Excel 文件失败", zap.Error(err))
		return err
	}

	rows := f.GetRows("Sheet1")
	if len(rows) == 0 {
		zap.L().Error("在Excel文件中找不到行")
		return errors.New("在Excel文件中找不到行")
	}

	for _, row := range rows {
		if len(row) < 2 {
			zap.L().Error("Excel文件中数据不完整", zap.Strings("row", row))
			return errors.New("Excel文件中数据不完整")
		}

		student := models.Student{
			Name:   row[0],
			Number: row[1],
			Status: "未签到", // 设置默认值
		}

		zap.L().Info("保存学生信息", zap.Any("student", student))
		if err := student.Save(); err != nil {
			zap.L().Error("保存学生信息失败", zap.Error(err))
			return err
		}
	}

	zap.L().Info("成功处理 Excel 文件", zap.String("filePath", filePath))
	return nil
}
