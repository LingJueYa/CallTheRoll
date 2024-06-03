package services

import (
	"CallTheRoll/api/models"
	"go.uber.org/zap"
)

// StudentStatus 学生的签到状态
type StudentStatus struct {
	Number string `json:"number"`
	Status string `json:"status"`
}

// UpdateStatus 接受学生学号和签到状态
func UpdateStatus(number, status string) error {
	zap.L().Info("更新签到状态", zap.String("number", number), zap.String("status", status))
	if err := models.UpdateStudentStatus(number, status); err != nil {
		zap.L().Error("更新学生签到状态失败", zap.String("number", number), zap.Error(err))
		return err
	}
	return nil
}

// GetStudents 获取所有学生信息
func GetStudents() ([]models.Student, error) {
	students, err := models.GetAllStudents()
	if err != nil {
		zap.L().Error("获取学生信息失败", zap.Error(err))
		return nil, err
	}
	return students, nil
}
