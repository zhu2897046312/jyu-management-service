package models

import (
	"jyu-service/utils"

	"gorm.io/gorm"
)

// `gorm:"not null;unique" json:"username"` 设置约束和json
// `gorm:"column:email_address;not null;unique;index" json:"email"` 设置字段名和索引

// 课程信息
type CourseInformation struct {
	CourseCode          string  `gorm:"not null;" json:"course_code"`             // 课程代码
	AcademicYear        string  `json:"academic_year"`        // 学年
	Semester            int     `json:"semester"`             // 学期
	CourseName          string  `json:"course_name"`          // 课程名称
	CommencementAcademy string  `json:"commencement_academy"` // 开课学院
	CourseType          int     `json:"course_type"`          // 课程类别
	CourseNature        int     `json:"course_nature"`        // 课程性质
	Credits             float32 `json:"credits"`              // 学分
	ClassName           string  `json:"class_name"`           // 教学班名称
	TeacherName         string  `json:"teacher_name"`         // 教师名称
	ClassTime           string  `json:"class_time"`           // 上课时间
	ClassAddress        string  `json:"class_address"`        // 上课地点
}


// 成绩信息
type GradeInformation struct {
	CourseCode   string  `gorm:"not null;"` // 课程代码
	Credits      float32 // 学分
	AcademicYear string  // 学年
	Semester     int     // 学期
	CourseName   string  // 课程名称
	Grade        float32 // 成绩
	GradePoints  float32 // 绩点
	CourseNature int     // 课程性质
	ExamNature   int     // 考试性质
	GradeRemark  string  // 成绩备注
}

// 课程性质
const (
	MajorsCompulsory           int = 1 // 专业必修
	MajorsElectives            int = 2 // 专业选修
	GeneralEducationCompulsory int = 3 // 通识必修
	GeneralEducationElectives  int = 4 // 通识选修
	CareerCompulsory           int = 5 // 职业必修
	CareerElectives            int = 6 // 职业选修
)

// 课程类别
const (
	MajorsCourse int = 1 // 专业课
)

// 考试性质
const (
	Normal    int = 1 // 正常考试
	Retakes_1 int = 2 // 补考1
	Retakes_2 int = 3 // 补考2
	Retakes_3 int = 4 // 补考3
	Retakes_4 int = 5 // 补考4
	Retakes_5 int = 5 // 补考5
)

func init() {
	err := utils.DB_MySQL.AutoMigrate(&CourseInformation{})
	if err != nil {
		panic(err)
	}
	err = utils.DB_MySQL.AutoMigrate(&GradeInformation{})
	if err != nil {
		panic(err)
	}
}

// 课程信息
func (u *CourseInformation) Insert(employee *CourseInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&CourseInformation{}).Create(employee)
}

func (u *CourseInformation) FindByCourseCode(user *CourseInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&CourseInformation{}).Where("course_code = ?", user.CourseCode).Find(&user)
}

func (u *CourseInformation) Update(employee *CourseInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&CourseInformation{}).Where("course_code = ?", employee.CourseCode).Updates(&employee)
}

func (u *CourseInformation) Delete(employee *CourseInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&CourseInformation{}).Where("course_code= ?", employee.CourseCode).Delete(&CourseInformation{})
}

func (u *CourseInformation) GetAll() ([]CourseInformation, *gorm.DB) {
	var borrowers []CourseInformation
	return borrowers, utils.DB_MySQL.Model(&CourseInformation{}).Find(&borrowers)
}

func (u *CourseInformation) PageQuery(page int, pageSize int) ([]CourseInformation, *gorm.DB) {
	employees := make([]CourseInformation, 0)
	var total int64

	utils.DB_MySQL.Model(&CourseInformation{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&CourseInformation{}).Limit(pageSize).Offset(offset).Find(&employees)
	return employees, query
}

// 成绩信息

func (u *GradeInformation) Insert(employee *GradeInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&GradeInformation{}).Create(employee)
}

func (u *GradeInformation) FindByCourseCode(user *GradeInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&GradeInformation{}).Where("course_code = ?", user.CourseCode).Find(&user)
}

func (u *GradeInformation) Update(employee *GradeInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&GradeInformation{}).Where("course_code = ?", employee.CourseCode).Updates(&employee)
}

func (u *GradeInformation) Delete(employee *GradeInformation) *gorm.DB {
	return utils.DB_MySQL.Model(&GradeInformation{}).Where("course_code= ?", employee.CourseCode).Delete(&GradeInformation{})
}

func (u *GradeInformation) GetAll() ([]GradeInformation, *gorm.DB) {
	var borrowers []GradeInformation
	return borrowers, utils.DB_MySQL.Model(&GradeInformation{}).Find(&borrowers)
}

func (u *GradeInformation) PageQuery(page int, pageSize int) ([]GradeInformation, *gorm.DB) {
	employees := make([]GradeInformation, 0)
	var total int64

	utils.DB_MySQL.Model(&GradeInformation{}).Count(&total)
	offset := (page - 1) * pageSize
	query := utils.DB_MySQL.Model(&GradeInformation{}).Limit(pageSize).Offset(offset).Find(&employees)
	return employees, query
}
