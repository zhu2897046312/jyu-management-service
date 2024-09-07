package models

import (
	"errors"
	"jyu-service/utils"

	"gorm.io/gorm"
)

type UserCourse struct {
	Account    string `gorm:"not null" json:"account"`    // 学号
	CourseCode string `gorm:"not null" json:"courseCode"` // 课程代码
	Status     int    `gorm:"not null" json:"status"`     // 选课状态码
}

func init() {
	err := utils.DB_MySQL.AutoMigrate(&UserCourse{})
	if err != nil {
		panic(err)
	}
}

func (table UserCourse) TableNanme() string {
	return "user_course"
}

func (u *UserCourse) Insert(employee *UserCourse) *gorm.DB {
	return utils.DB_MySQL.Model(&UserCourse{}).Create(employee)
}

func (u *UserCourse) FindByAccount(user *UserCourse) *gorm.DB {
	return utils.DB_MySQL.Model(&UserCourse{}).Where("account = ?", user.Account).Find(&user)
}

func (u *UserCourse) Update(employee *UserCourse) *gorm.DB {
	return utils.DB_MySQL.Model(&UserCourse{}).Where("account = ?", employee.Account).Updates(&employee)
}

func (u *UserCourse) Delete(employee *UserCourse) *gorm.DB {
	return utils.DB_MySQL.Model(&UserCourse{}).Where("account= ?", employee.Account).Delete(&UserCourse{})
}

func (u *UserCourse) GetAll() ([]UserCourse, *gorm.DB) {
	var borrowers []UserCourse
	return borrowers, utils.DB_MySQL.Model(&UserCourse{}).Find(&borrowers)
}

func (u *UserCourse) PageQuery(page int, pageSize int) ([]UserCourse, *gorm.DB) {
	employees := make([]UserCourse, 0)
	var total int64

	utils.DB_MySQL.Model(&UserCourse{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&UserCourse{}).Limit(pageSize).Offset(offset).Find(&employees)
	return employees, query
}

// 插入选课信息到 UserCourse 中间表
func (u *UserCourse) EnrollCourse(account string, courseCode string) error {
	// 检查是否已选过该课程
	var existingEnrollment UserCourse
	if err := utils.DB_MySQL.Model(&UserCourse{}).Where("account = ? AND course_code = ?", account, courseCode).First(&existingEnrollment).Error; err == nil {
		return errors.New("该课程已被选过")
	}

	// 如果没有选过该课程，插入新的选课记录
	enrollment := UserCourse{
		Account:    account,
		CourseCode: courseCode,
	}

	// 插入到数据库
	if err := utils.DB_MySQL.Model(&UserCourse{}).Create(&enrollment).Error; err != nil {
		return err
	}

	return nil
}
