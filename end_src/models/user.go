package models

import (
	"jyu-service/utils"

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
	Account              string `gorm:"primaryKey;not null;" json:"account"`   // 学号
	Name                 string `gorm:"not null" json:"name"`                  // 姓名
	Sex                  int    `gorm:"not null" json:"sex"`                   // 性别
	IdentificationNumber string `gorm:"not null" json:"identification_number"` // 身份证号
	Birthday             string `gorm:"not null" json:"birthday"`              // 出生日期
	EthnicGroup          string `gorm:"not null" json:"ethnic_group"`          // 民族
	IdentificationType   string `gorm:"not null" json:"identification_type"`   // 证件类型
}

// 学籍信息
type StudentStatusInformation struct {
	Account            string `gorm:"primaryKey;not null;" json:"account"` // 学号
	Grade              string `gorm:"not null"`                            // 年级
	AcademyName        string `gorm:"not null"`                            // 学院名称
	ClassName          string `gorm:"not null"`                            // 班级名称
	ProfessionalName   string `gorm:"not null"`                            // 专业名称
	Status             string `gorm:"not null"`                            // 学籍状态
	IsInSchool         bool   `gorm:"not null"`                            // 是否在校
	RegistrationStatus string `gorm:"not null"`                            // 报到注册状态
	EducationalLevel   string `gorm:"not null"`                            // 学历层次
	CultivationMethod  string `gorm:"not null"`                            // 培养方式
	CultivationLevel   int    `gorm:"not null"`                            // 培养层次
	StudentType        int    `gorm:"not null"`                            // 学生类别
	CheckInTime        string `gorm:"not null"`                            // 报到时间
	RegistrationTime   string `gorm:"not null"`                            // 注册时间
}

// 联系方式
type ContactInformation struct {
	Account               string `gorm:"primaryKey;not null;"` // 学号
	CorrespondenceAddress string `gorm:"not null"`             // 通讯地址
	Phone                 string `gorm:"not null"`             // 手机号码
	Email                 string // 电子邮箱
	Landline              string // 固定电话
	PostCode              string // 邮政编码
}

func init() {
	err := utils.DB_MySQL.AutoMigrate(&UserAccount{})
	if err != nil {
		panic(err)
	}
	err = utils.DB_MySQL.AutoMigrate(&UserBasicInformation{})
	if err != nil {
		panic(err)
	}
	err = utils.DB_MySQL.AutoMigrate(&StudentStatusInformation{})
	if err != nil {
		panic(err)
	}
	err = utils.DB_MySQL.AutoMigrate(&ContactInformation{})
	if err != nil {
		panic(err)
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

func (u *StudentStatusInformation) FindByID(user *StudentStatusInformation) *gorm.DB {
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

func (u *ContactInformation) FindByID(user *ContactInformation) *gorm.DB {
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
