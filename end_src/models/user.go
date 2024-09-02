package models


import (
	"jyu-service/utils"
	"gorm.io/gorm"
)

// 基本信息
type UserBasicInformation struct {
	Account              string `gorm:"not null;unique"` // 学号
	Password             string `gorm:"not null"`        // 密码
	Name                 string // 姓名
	Sex                  int    // 性别
	IdentificationNumber string // 身份证号
	Birthday             string // 出生日期
	EthnicGroup          string // 民族
	IdentificationType   string // 证件类型
	ChatType             int    // 账户类型 学生 、 教师 、 管理员
}

// 学籍信息
type StudentStatusInformation struct {
	Account            string `gorm:"not null;unique"` // 学号
	Grade              string // 年级
	AcademyName        string // 学院名称
	ClassName          string // 班级名称
	ProfessionalName   string // 专业名称
	Status             string // 学籍状态
	IsInSchool         bool   // 是否在校
	RegistrationStatus string // 报到注册状态
	EducationalLevel   string // 学历层次
	CultivationMethod  string // 培养方式
	CultivationLevel   int    // 培养层次
	StudentType        int    // 学生类别
	CheckInTime        string // 报到时间
	RegistrationTime   string // 注册时间
}

// 联系方式
type ContactInformation struct {
	Account               string `gorm:"not null;unique"` // 学号
	Email                 string // 电子邮箱
	Phone                 string // 手机号码
	Landline              string // 固定电话
	PostCode              string // 邮政编码
	CorrespondenceAddress string // 通讯地址
}

func (table UserBasicInformation) TableNanme() string {
	return "user_basic_information"
}

func (table *UserBasicInformation) Find(account string) (*UserBasicInformation,*gorm.DB) {
	var user *UserBasicInformation
	return user, utils.DB_MySQL.Model(&UserBasicInformation{}).Where("account = ?", account).Find(user)
}

func (table *UserBasicInformation)Create(user UserBasicInformation) (*gorm.DB){
	return utils.DB_MySQL.Model(&UserBasicInformation{}).Create(&user)
}

func (table *UserBasicInformation)Delete(account string) (*gorm.DB){
	return utils.DB_MySQL.Model(&UserBasicInformation{}).Where("account = ?", account).Delete(&UserBasicInformation{})
}

func  (table *UserBasicInformation)Update(user UserBasicInformation) *gorm.DB {
    // 执行更新操作
    return utils.DB_MySQL.Where("account = ?", user.Account).Model(user).Updates(UserBasicInformation{
		Password: user.Password,
		Name: user.Name,
		Sex: user.Sex,
        IdentificationNumber: user.IdentificationNumber,
        Birthday: user.Birthday,
        EthnicGroup: user.EthnicGroup,
        IdentificationType: user.IdentificationType,
	})
}

func (table StudentStatusInformation) TableNanme() string {
	return "student_status_information"
}

func (table ContactInformation) TableNanme() string {
	return "contact_information"
}


