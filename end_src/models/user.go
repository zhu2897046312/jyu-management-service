package models

import (
	"jyu-service/utils"
	"log"
	"strconv"
	"time"
	"fmt"
	"gorm.io/gorm"
)

// 账号
type UserAccount struct {
	Account  string `gorm:"primaryKey;not null;" json:"account"` // 学号
	Password string `gorm:"not null" json:"password"`            // 密码
	ChatType int    `gorm:"not null" json:"chat_type"`           // 账户类型 学生 、 教师 、 管理员
}

// 基本信息
type UserBasicInformation struct {
	Account              string `gorm:"primaryKey;not null;" json:"account"` // 学号
	Name                 string `json:"name"`                                // 姓名
	Sex                  int    `json:"sex"`                                 // 性别
	IdentificationNumber string `json:"identification_number"`               // 身份证号
	Birthday             string `json:"birthday"`                            // 出生日期
	EthnicGroup          string `json:"ethnic_group"`                        // 民族
	IdentificationType   string `json:"identification_type"`                 // 证件类型
	OldName              string `json:"old_name"`                            // 曾用名
	PoliticalOutlook     string `json:"political_outlook"`                   // 政治面貌
	EnrollmentDates      string `json:"enrollment_dates"`                    // 入学日期
}

// 学籍信息
type StudentStatusInformation struct {
	Account            string `gorm:"primaryKey;not null;" json:"account"` // 学号
	AcademicYear       string `json:"academic_year"`                       // 年级
	AcademyName        string `gorm:"not null" json:"academy_name"`        // 学院名称
	ClassName          string `gorm:"not null" json:"class_name"`          // 班级名称
	ProfessionalName   string `gorm:"not null" json:"professional_name"`   // 专业名称
	Status             string `gorm:"not null" json:"status"`              // 学籍状态
	IsInSchool         int    `gorm:"not null" json:"is_in_School"`        // 是否在校
	RegistrationStatus string `gorm:"not null" json:"registration_status"` // 报到注册状态
	EducationalLevel   string `gorm:"not null" json:"educational_level"`   // 学历层次
	CultivationMethod  string `gorm:"not null" json:"cultivation_method"`  // 培养方式
	CultivationLevel   int    `gorm:"not null" json:"cultivation_level"`   // 培养层次
	StudentType        int    `gorm:"not null" json:"student_type"`        // 学生类别
	CheckInTime        string `gorm:"not null" json:"check_in_time"`       // 报到时间
	RegistrationTime   string `gorm:"not null" json:"registration_time"`   // 注册时间
	Academic           int    `gorm:"not null" json:"academic"`            // 学制
}

// 联系方式
type ContactInformation struct {
	Account               string `gorm:"primaryKey;not null;" json:"account"`    // 学号
	CorrespondenceAddress string `gorm:"not null" json:"correspondence_address"` // 通讯地址Correspondence address
	Phone                 string `gorm:"not null" json:"phone"`                  // 手机号码
	Email                 string `gorm:"not null" json:"email"`                  // 电子邮箱
	Landline              string `gorm:"not null" json:"landline"`               // 固定电话
	PostCode              string `gorm:"not null" json:"post_code"`              // 邮政编码
	HomeAddress           string `gorm:"not null" json:"home_address"`           // 家庭地址
}

const (
	Administrator int = 0
	Teacher       int = 2
	Student       int = 1
)

func init() {
	// 自动迁移 UserAccount
	if err := utils.DB_MySQL.AutoMigrate(&UserAccount{}); err != nil {
		log.Fatalf("Failed to migrate UserAccount: %v", err)
	} else {
		log.Println("UserAccount table migrated successfully")
	}

	// 自动迁移 UserBasicInformation
	if err := utils.DB_MySQL.AutoMigrate(&UserBasicInformation{}); err != nil {
		log.Fatalf("Failed to migrate UserBasicInformation: %v", err)
	} else {
		log.Println("UserBasicInformation table migrated successfully")
	}

	// 自动迁移 StudentStatusInformation
	if err := utils.DB_MySQL.AutoMigrate(&StudentStatusInformation{}); err != nil {
		log.Fatalf("Failed to migrate StudentStatusInformation: %v", err)
	} else {
		log.Println("StudentStatusInformation table migrated successfully")
	}

	// 自动迁移 ContactInformation
	if err := utils.DB_MySQL.AutoMigrate(&ContactInformation{}); err != nil {
		log.Fatalf("Failed to migrate ContactInformation: %v", err)
	} else {
		log.Println("ContactInformation table migrated successfully")
	}
}

func (table UserAccount) TableNanme() string {
	return "user_account"
}

func (table UserBasicInformation) TableNanme() string {
	return "user_basic_information"
}

func (table StudentStatusInformation) TableNanme() string {
	return "student_status_information"
}

func (table ContactInformation) TableNanme() string {
	return "contact_information"
}

// 账号
// GenerateAccount 生成账号逻辑
func GenerateAccount(chatType int) string {
	// 1. 获取当前年份后两位
	currentYear := time.Now().Year()
	yearSuffix := strconv.Itoa(currentYear)[2:] // 取后两位

	// 2. 根据 chatType 确定账号类型标志
	var accountType string
	switch chatType {
	case Administrator:
		accountType = "0"
	case Student:
		accountType = "1"
	case Teacher:
		accountType = "2"
	default:
		accountType = "1" // 默认学生
	}

	// 3. 查找当前最大账号并自增
	var lastAccount UserAccount
	err := utils.DB_MySQL.Model(&UserAccount{}).
		Where("account LIKE ?", fmt.Sprintf("%s%s%%", yearSuffix, accountType)).
		Order("account DESC").First(&lastAccount).Error

	var nextID int
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("Error fetching last account:", err)
		return ""
	}

	if err == gorm.ErrRecordNotFound {
		// 如果没有找到，表示是今年第一个该类型账号，初始值为00001
		nextID = 1
	} else {
		// 取出账号最后五位并加1
		lastIDStr := lastAccount.Account[len(lastAccount.Account)-5:]
		lastID, _ := strconv.Atoi(lastIDStr)
		nextID = lastID + 1
	}

	// 4. 将 nextID 格式化为五位数，并拼接成完整账号
	accountSuffix := fmt.Sprintf("%05d", nextID) // 保证自增部分是5位数
	newAccount := yearSuffix + accountType + accountSuffix

	return newAccount
}

func (u *UserAccount) Insert_auto() (string ,*gorm.DB) {
	//2. 生成账号
	u.Account = GenerateAccount(u.ChatType)
	// 插入数据库
	return u.Account ,utils.DB_MySQL.Model(&UserAccount{}).Create(&u)
}

func (u *UserAccount) Insert(employee *UserAccount) *gorm.DB {
	return utils.DB_MySQL.Model(&UserAccount{}).Create(employee)
}

func (u *UserAccount) FindByAccount(user *UserAccount) *gorm.DB {
	return utils.DB_MySQL.Model(&UserAccount{}).Where("account = ?", user.Account).Find(&user)
}

func (u *UserAccount) Update(employee *UserAccount) *gorm.DB {
	return utils.DB_MySQL.Model(&UserAccount{}).Where("account = ?", employee.Account).Updates(&employee)
}

func (u *UserAccount) Delete(employee *UserAccount) *gorm.DB {
	return utils.DB_MySQL.Model(&UserAccount{}).Where("account= ?", employee.Account).Delete(&UserAccount{})
}

func (u *UserAccount) GetAll() ([]UserAccount, *gorm.DB) {
	var borrowers []UserAccount
	return borrowers, utils.DB_MySQL.Model(&UserAccount{}).Find(&borrowers)
}

func (u *UserAccount) PageQuery(page int, pageSize int) ([]UserAccount, *gorm.DB) {
	employees := make([]UserAccount, 0)
	var total int64

	utils.DB_MySQL.Model(&UserAccount{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&UserAccount{}).Limit(pageSize).Offset(offset).Find(&employees)
	return employees, query
}

// UserBasicInformation	基本信息
func (u *UserBasicInformation) Insert(employee *UserBasicInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&UserBasicInformation{}).Create(employee)
}

func (u *UserBasicInformation) FindByAccount(user *UserBasicInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&UserBasicInformation{}).Where("account = ?", user.Account).Find(&user)
}

func (u *UserBasicInformation) Update(employee *UserBasicInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&UserBasicInformation{}).Where("account = ?", employee.Account).Updates(&employee)
}

func (u *UserBasicInformation) Delete(employee *UserBasicInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&UserBasicInformation{}).Where("account= ?", employee.Account).Delete(&UserBasicInformation{})
}

func (u *UserBasicInformation) GetAll() ([]UserBasicInformation, *gorm.DB) {
	var borrowers []UserBasicInformation
	return borrowers, utils.DB_MySQL.Model(&UserBasicInformation{}).Find(&borrowers)
}

func (u *UserBasicInformation) PageQuery(page int, pageSize int) ([]UserBasicInformation, *gorm.DB) {
	employees := make([]UserBasicInformation, 0)
	var total int64

	utils.DB_MySQL.Model(&UserBasicInformation{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&UserBasicInformation{}).Limit(pageSize).Offset(offset).Find(&employees)
	return employees, query
}

// StudentStatusInformation	学籍信息
func (u *StudentStatusInformation) Insert(employee *StudentStatusInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&StudentStatusInformation{}).Create(employee)
}

func (u *StudentStatusInformation) FindByAccount(user *StudentStatusInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&StudentStatusInformation{}).Where("account = ?", user.Account).Find(&user)
}

func (u *StudentStatusInformation) Update(employee *StudentStatusInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&StudentStatusInformation{}).Where("account = ?", employee.Account).Updates(&employee)
}

func (u *StudentStatusInformation) Delete(employee *StudentStatusInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&StudentStatusInformation{}).Where("account= ?", employee.Account).Delete(&StudentStatusInformation{})
}

func (u *StudentStatusInformation) GetAll() ([]StudentStatusInformation, *gorm.DB) {
	var borrowers []StudentStatusInformation
	return borrowers, utils.DB_MySQL.Model(&StudentStatusInformation{}).Find(&borrowers)
}

func (u *StudentStatusInformation) PageQuery(page int, pageSize int) ([]StudentStatusInformation, *gorm.DB) {
	employees := make([]StudentStatusInformation, 0)
	var total int64

	utils.DB_MySQL.Model(&StudentStatusInformation{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&StudentStatusInformation{}).Limit(pageSize).Offset(offset).Find(&employees)
	return employees, query
}

// ContactInformation	联系方式
func (u *ContactInformation) Insert(employee *ContactInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&ContactInformation{}).Create(employee)
}

func (u *ContactInformation) FindByAccount(user *ContactInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&ContactInformation{}).Where("account = ?", user.Account).Find(&user)
}

func (u *ContactInformation) Update(employee *ContactInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&ContactInformation{}).Where("account = ?", employee.Account).Updates(&employee)
}

func (u *ContactInformation) Delete(employee *ContactInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&ContactInformation{}).Where("account= ?", employee.Account).Delete(&ContactInformation{})
}

func (u *ContactInformation) GetAll() ([]ContactInformation, *gorm.DB) {
	var borrowers []ContactInformation
	return borrowers, utils.DB_MySQL.Model(&ContactInformation{}).Find(&borrowers)
}

func (u *ContactInformation) PageQuery(page int, pageSize int) ([]ContactInformation, *gorm.DB) {
	employees := make([]ContactInformation, 0)
	var total int64

	utils.DB_MySQL.Model(&ContactInformation{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&ContactInformation{}).Limit(pageSize).Offset(offset).Find(&employees)
	return employees, query
}
