package models

// `gorm:"not null;unique" json:"username"` 设置约束和json
// `gorm:"column:email_address;not null;unique;index" json:"email"` 设置字段名和索引

// 选课信息
type CourseInformation struct {
	CourseCode          string  `gorm:"not null;unique"` // 课程代码
	AcademicYear        string  // 学年
	Semester            int     // 学期
	CourseName          string  // 课程名称
	CommencementAcademy string  // 开课学院
	CourseType          int     // 课程类别
	CourseNature        int     // 课程性质
	Credits             float32 // 学分
	ClassName           string  // 教学班名称
	TeacherName         string  // 教师名称
	ClassTime           string  // 上课时间
	ClassAddress        string  // 上课地点
}

// 成绩信息
type GradeInformation struct {
	CourseCode   string  `gorm:"not null;unique"` // 课程代码
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
