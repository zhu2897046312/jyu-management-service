package models

import (
	"fmt"
	"jyu-service/utils"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// `gorm:"not null;unique" json:"username"` 设置约束和json
// `gorm:"column:email_address;not null;unique;index" json:"email"` 设置字段名和索引

// 课程信息
type CourseInformation struct {
	CourseCode          string  `gorm:"primaryKey;not null;" json:"course_code"` // 课程代码
	Account             string  `json:"account"`                                 // 教师账号
	TeacherName         string  `json:"teacher_name"`                            // 教师名称
	AcademicYear        string  `json:"academic_year"`                           // 年级
	Semester            int     `json:"semester"`                                // 学期
	CourseName          string  `json:"course_name"`                             // 课程名称
	CommencementAcademy string  `json:"commencement_academy"`                    // 开课学院
	CourseAffiliation   string  `json:"course_affiliation"`                      // 课程归属
	CourseType          int     `json:"course_type"`                             // 课程类别
	CourseNature        int     `json:"course_nature"`                           // 课程性质
	Credits             float32 `json:"credits"`                                 // 学分
	ClassName           string  `json:"class_name"`                              // 教学班名称
	ClassTime           string  `json:"class_time"`                              // 上课时间（节次:星期:持续周次）
	ClassAddress        string  `json:"class_address"`                           // 上课地点
	MaxStudentNumber    int     `json:"max_student_number"`                      // 人数
	ChoosedNumber       int     `json:"choosed_number"`                          // 已选人数
	TeachingMode        int     `json:"teaching_mode"`                           // 教学模式
}

// 成绩信息
type GradeInformation struct {
	CourseCode   string  `gorm:"primaryKey;not null;" json:"course_code"` // 课程代码
	Credits      float32 `json:"credits"`                                 // 学分
	AcademicYear string  `json:"academic_year"`                           // 学年
	Semester     int     `json:"semester"`                                // 学期
	CourseName   string  `json:"course_name"`                             // 课程名称
	Grade        float32 `json:"grade "`                                  // 成绩
	GradePoints  float32 `json:"grade_points"`                            // 绩点
	CourseNature int     `json:"course_nature"`                           // 课程性质
	ExamNature   int     `json:"exam_nature"`                             // 考试性质
	GradeRemark  string  `json:"grade_remark"`                            // 成绩备注
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

const (
	Online  int = 1 //线上
	Offline int = 2 //线下
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

func (u *CourseInformation) FindByCourseCode() *gorm.DB {
	return utils.DB_MySQL.Model(&CourseInformation{}).Where("course_code = ?", u.CourseCode).Find(&u)
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

func (u *CourseInformation) DynamicQuery(conditions map[string]interface{}) ([]CourseInformation, error) {
	var results []CourseInformation
	query := utils.DB_MySQL.Model(&CourseInformation{})

	// Build the query based on conditions
	for key, value := range conditions {
		switch key {
		case "course_code":
			query = query.Where("course_code = ?", value)
		case "academic_year":
			query = query.Where("academic_year = ?", value)
		case "semester":
			query = query.Where("semester = ?", value)
		case "course_name":
			query = query.Where("course_name LIKE ?", "%"+value.(string)+"%")
		case "commencement_academy":
			query = query.Where("commencement_academy = ?", value)
		case "course_type":
			query = query.Where("course_type = ?", value)
		case "course_nature":
			query = query.Where("course_nature = ?", value)
		case "credits":
			query = query.Where("credits = ?", value)
		case "class_name":
			query = query.Where("class_name LIKE ?", "%"+value.(string)+"%")
		case "teacher_name":
			query = query.Where("teacher_name LIKE ?", "%"+value.(string)+"%")
		case "class_time":
			query = query.Where("class_time = ?", value)
		case "class_address":
			query = query.Where("class_address LIKE ?", "%"+value.(string)+"%")
		case "max_student_number":
			query = query.Where("max_student_number = ?", value)
		case "choosed_number":
			query = query.Where("choosed_number = ?", value)
		case "teaching_mode":
			query = query.Where("teaching_mode = ?", value)
		default:
			return nil, fmt.Errorf("unsupported query condition: %s", key)
		}
	}

	// Execute the query
	err := query.Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

// 获取课程的已选人数
func (course *CourseInformation) GetCourseChoosedNumber() (int, error) {
	// 构建 Redis 中存储已选人数的键
	choosedNumberKey := "course:" + course.CourseCode + ":choosed_number"
	maxStudentNumberKey := "course:" + course.CourseCode + ":max_student_number"

	// 尝试从 Redis 中获取已选人数
	choosedNumber, err := utils.DB_Redis.Get(utils.Redis_Context, choosedNumberKey).Int()
	if err == redis.Nil {
		// 如果 Redis 中没有相关记录，从 MySQL 获取
		if err := utils.DB_MySQL.Where("course_code = ?", course.CourseCode).First(course).Error; err != nil {
			return 0, fmt.Errorf("从 MySQL 获取课程数据失败: %v", err)
		}

		// 将 MySQL 中的已选人数缓存到 Redis
		choosedNumber = course.ChoosedNumber
		maxStudentNumber := course.MaxStudentNumber

		err := utils.DB_Redis.Set(utils.Redis_Context, choosedNumberKey, choosedNumber, 0).Err()
		utils.DB_Redis.Set(utils.Redis_Context, maxStudentNumberKey, maxStudentNumber, 0)
		if err != nil {
			return 0, fmt.Errorf("缓存课程已选人数到 Redis 失败: %v", err)
		}
	} else if err != nil {
		return 0, fmt.Errorf("从 Redis 获取已选人数失败: %v", err)
	}

	// 返回已选人数
	return choosedNumber, nil
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
