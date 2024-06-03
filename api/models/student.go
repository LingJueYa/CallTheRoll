package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

// Student 数据模型
type Student struct {
	Name   string
	Number string
	Status string
}

// Save 将学生信息保存在数据库
func (s *Student) Save() error {
	// 打开数据库文件
	db, err := sql.Open("sqlite3", "./students.db")
	if err != nil {
		zap.L().Error("打开数据库失败", zap.Error(err))
		return err
	}
	defer db.Close()
	// 如果学生状态为空，则默认设置为“未签到”
	if s.Status == "" {
		s.Status = "未签到"
	}
	// 执行插入
	_, err = db.Exec("INSERT INTO students (name, number, status) VALUES (?, ?, ?)", s.Name, s.Number, s.Status)
	if err != nil {
		zap.L().Error("插入学生信息失败", zap.Error(err))
		return err
	}
	zap.L().Info("存入学生状态成功", zap.String("name", s.Name), zap.String("number", s.Number), zap.String("status", s.Status))
	return nil
}

// UpdateStudentStatus 更新学生签到状态
func UpdateStudentStatus(number, status string) error {
	db, err := sql.Open("sqlite3", "./students.db")
	if err != nil {
		zap.L().Error("打开数据库失败", zap.Error(err))
		return err
	}
	defer db.Close()
	// 执行更新
	_, err = db.Exec("UPDATE students SET status = ? WHERE number = ?", status, number)
	if err != nil {
		zap.L().Error("更新学生状态失败", zap.Error(err))
		return err
	}

	zap.L().Info("更新学生状态成功", zap.String("number", number), zap.String("status", status))
	return nil
}

// GetAllStudents 获取所有学生信息
func GetAllStudents() ([]Student, error) {
	db, err := sql.Open("sqlite3", "./students.db")
	if err != nil {
		zap.L().Error("打开数据库失败", zap.Error(err))
		return nil, err
	}
	defer db.Close()
	// 执行查询
	rows, err := db.Query("SELECT name, number, IFNULL(status, '未签到') FROM students")
	if err != nil {
		zap.L().Error("查询学生信息失败", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.Name, &s.Number, &s.Status); err != nil {
			zap.L().Error("遍历学生信息失败", zap.Error(err))
			return nil, err
		}
		// 插入到切片
		students = append(students, s)
	}
	// 返回学生列表
	zap.L().Info("查询成功", zap.Int("count", len(students)))
	return students, nil
}

// ClearStudents 清空数据库
func ClearStudents() error {
	db, err := sql.Open("sqlite3", "./students.db")
	if err != nil {
		zap.L().Error("打开数据库失败", zap.Error(err))
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM students")
	if err != nil {
		zap.L().Error("清空数据库失败", zap.Error(err))
		return err
	}

	zap.L().Info("所有学生数据已清空")
	return nil
}
