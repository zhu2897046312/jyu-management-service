package models

import (
	"errors"
	"jyu-service/utils"
	"strings"
	"fmt"
	"time"
	"gorm.io/gorm"
)

type UserCourse struct {
	Account    string `gorm:"not null" json:"account"`    // 学号
	CourseCode string `gorm:"not null" json:"course_code"` // 课程代码
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

func (u *UserCourse) GetByAccount(account string) ([]UserCourse, *gorm.DB) {
	var borrowers []UserCourse
	return borrowers, utils.DB_MySQL.Model(&UserCourse{}).Where("account= ?", account).Find(&borrowers)
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

// 退选
func (u *UserCourse)UnenrollCourse() *gorm.DB{
	// 先检查该课程是否存在
	return utils.DB_MySQL.Model(&UserCourse{}).Where("account = ? AND course_code = ?", u.Account,u.CourseCode).Delete(&UserCourse{})
}

// 提取账户信息的函数
func extractAccountFromKey(userKey string) string {
	// 从 Redis 键中提取账户信息的逻辑
	// 假设键的格式是 "enrollments:<account>"
	parts := strings.Split(userKey, ":")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

func UpdateCourseEnrollmentsFromRedis() {
	ticker := time.NewTicker(1 * time.Minute) // 每分钟执行一次
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 获取所有用户的选课记录
			users, err := utils.DB_Redis.Keys(utils.Redis_Context, "enrollments:*").Result()
			if err != nil {
				// 处理错误
				fmt.Println("获取用户选课记录失败:", err)
				continue
			}

			for _, userKey := range users {
				account := extractAccountFromKey(userKey) // 提取账户信息
				courseCodes, err := utils.DB_Redis.SMembers(utils.Redis_Context, userKey).Result()
				if err != nil {
					// 处理错误
					fmt.Println("获取用户选课记录失败:", err)
					continue
				}

				for _, courseCode := range courseCodes {
					// 检查记录是否已存在于 MySQL
					var existingEnrollment UserCourse
					err := utils.DB_MySQL.Where("account = ? AND course_code = ?", account, courseCode).First(&existingEnrollment).Error
					if err != nil && err != gorm.ErrRecordNotFound {
						// 处理错误
						fmt.Println("查询选课记录失败:", err)
						continue
					}

					// 如果记录不存在，则插入
					if err == gorm.ErrRecordNotFound {
						enrollment := UserCourse{Account: account, CourseCode: courseCode}
						if err := utils.DB_MySQL.Create(&enrollment).Error; err != nil {
							// 处理错误
							fmt.Println("插入选课记录失败:", err)
							continue
						}
					}

					// 更新课程信息表中的已选人数
					courseKey := "course:" + courseCode
					choosedNumberKey := courseKey + ":choosed_number"

					// 从 Redis 获取当前已选人数
					choosedNumber, err := utils.DB_Redis.Get(utils.Redis_Context, choosedNumberKey).Int()
					if err != nil {
						// 处理错误
						fmt.Println("获取已选人数失败:", err)
						continue
					}

					// 更新 MySQL 中的课程信息
					var course CourseInformation
					err = utils.DB_MySQL.Where("course_code = ?", courseCode).First(&course).Error
					if err != nil {
						// 处理错误
						fmt.Println("查询课程信息失败:", err)
						continue
					}

					// 更新已选人数
					course.ChoosedNumber = choosedNumber
					if err := utils.DB_MySQL.Save(&course).Error; err != nil {
						// 处理错误
						fmt.Println("更新课程信息失败:", err)
					}
				}
			}
		}
	}
}