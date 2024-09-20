package service

import (
	"jyu-service/models"
	"jyu-service/utils"
	"net/http"
	"strconv"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// 查询并生成 Excel 文档
func GenerateExcel(c *gin.Context) {
	var courses []models.CourseInformation
	var userCourses []models.UserCourse

	// 查询数据
	if err := utils.DB_MySQL.Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := utils.DB_MySQL.Find(&userCourses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 创建 Excel 文件
	f := excelize.NewFile()
	sheetName := "Sheet1"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	f.SetActiveSheet(index)

	// 添加表头
	f.SetCellValue(sheetName, "A1", "学号")
	f.SetCellValue(sheetName, "B1", "学年")
	f.SetCellValue(sheetName, "C1", "学期")
	f.SetCellValue(sheetName, "D1", "课程代码")
	f.SetCellValue(sheetName, "E1", "课程名字")
	f.SetCellValue(sheetName, "F1", "成绩")
	f.SetCellValue(sheetName, "G1", "绩点")

	// 添加数据
	row := 2
	for _, uc := range userCourses {
		if uc.CourseGrade == "" { // 跳过成绩为空的记录
			continue
		}

		for _, c := range courses {
			if c.CourseCode == uc.CourseCode {
				f.SetCellValue(sheetName, "A"+strconv.Itoa(row), uc.Account)
				f.SetCellValue(sheetName, "B"+strconv.Itoa(row), c.AcademicYear)
				f.SetCellValue(sheetName, "C"+strconv.Itoa(row), c.Semester)
				f.SetCellValue(sheetName, "D"+strconv.Itoa(row), c.CourseCode)
				f.SetCellValue(sheetName, "E"+strconv.Itoa(row), c.CourseName)
				f.SetCellValue(sheetName, "F"+strconv.Itoa(row), uc.CourseGrade)
				f.SetCellValue(sheetName, "G"+strconv.Itoa(row), CalculateGradePoints(uc.CourseGrade,c.Credits))
				row++
			}
		}
	}

	// 设置 Excel 输出流
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=student_grades.xlsx")
	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}


// 生成固定格式的 Excel 文件
func GenerateExcelTemplate(c *gin.Context) {
	f := excelize.NewFile()

	// 创建“Students” Sheet 并写入表头
	index, err := f.NewSheet("Students")
	if err!= nil {
        log.Fatal(err)
    }
	f.SetCellValue("Students", "A1", "Account")            // 学号
	f.SetCellValue("Students", "B1", "AcademicYear")       // 年级
	f.SetCellValue("Students", "C1", "AcademyName")        // 学院名称
	f.SetCellValue("Students", "D1", "ClassName")          // 班级名称
	f.SetCellValue("Students", "E1", "ProfessionalName")   // 专业名称
	f.SetCellValue("Students", "F1", "Status")             // 学籍状态
	f.SetCellValue("Students", "G1", "IsInSchool")         // 是否在校
	f.SetCellValue("Students", "H1", "RegistrationStatus") // 报到注册状态
	f.SetCellValue("Students", "I1", "EducationalLevel")   // 学历层次
	f.SetCellValue("Students", "J1", "CultivationMethod")  // 培养方式
	f.SetCellValue("Students", "K1", "CultivationLevel")   // 培养层次
	f.SetCellValue("Students", "L1", "StudentType")        // 学生类别
	f.SetCellValue("Students", "M1", "CheckInTime")        // 报到时间
	f.SetCellValue("Students", "N1", "RegistrationTime")   // 注册时间
	f.SetCellValue("Students", "O1", "Academic")           // 学制

	// 创建“Teachers” Sheet 并写入表头
	f.NewSheet("Teachers")
	f.SetCellValue("Teachers", "A1", "Account")            // 学号
	f.SetCellValue("Teachers", "B1", "Password")           // 密码
	f.SetCellValue("Teachers", "C1", "ChatType")           // 账户类型
	f.SetCellValue("Teachers", "D1", "CorrespondenceAddress") // 通讯地址
	f.SetCellValue("Teachers", "E1", "Phone")              // 手机号码
	f.SetCellValue("Teachers", "F1", "Email")              // 电子邮箱
	f.SetCellValue("Teachers", "G1", "Landline")           // 固定电话
	f.SetCellValue("Teachers", "H1", "PostCode")           // 邮政编码
	f.SetCellValue("Teachers", "I1", "HomeAddress")        // 家庭地址

	// 创建“Admins” Sheet 并写入表头
	f.NewSheet("Admins")
	f.SetCellValue("Admins", "A1", "Account")            // 学号
	f.SetCellValue("Admins", "B1", "Password")           // 密码
	f.SetCellValue("Admins", "C1", "ChatType")           // 账户类型
	f.SetCellValue("Admins", "D1", "CorrespondenceAddress") // 通讯地址
	f.SetCellValue("Admins", "E1", "Phone")              // 手机号码
	f.SetCellValue("Admins", "F1", "Email")              // 电子邮箱
	f.SetCellValue("Admins", "G1", "Landline")           // 固定电话
	f.SetCellValue("Admins", "H1", "PostCode")           // 邮政编码
	f.SetCellValue("Admins", "I1", "HomeAddress")        // 家庭地址

	// 设置默认Sheet为“Students”
	f.SetActiveSheet(index)

	// 设置响应头，返回文件流给前端
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=\"template.xlsx\"")
	c.Header("File-Name", "template.xlsx")
	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法生成 Excel 文件",
		})
		return
	}
}


// 添加、更新账号信息
// 批量导入Excel
func ImportFromExcel(filePath string) error {
	// 打开Excel文件
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("打开文件失败: %w", err)
	}
	defer f.Close()

	// 导入学生
	if err := importStudents(f); err != nil {
		log.Printf("学生导入失败: %v\n", err)
	}

	// 导入教师
	if err := importTeachers(f); err != nil {
		log.Printf("教师导入失败: %v\n", err)
	}

	// 导入管理员
	if err := importAdmins(f); err != nil {
		log.Printf("管理员导入失败: %v\n", err)
	}

	return nil
}

func importStudents(f *excelize.File) error {
	rows, err := f.GetRows("Students")
	if err != nil {
		return err
	}

	for _, row := range rows[1:] {
		student := models.StudentStatusInformation{
			Account:            row[0],
			AcademicYear:       row[1],
			AcademyName:        row[2],
			ClassName:          row[3],
			ProfessionalName:   row[4],
			Status:             row[5],
			IsInSchool:         parseBool(row[6]),
			RegistrationStatus: row[7],
			EducationalLevel:   row[8],
			CultivationMethod:  row[9],
			CultivationLevel:   parseInt(row[10]),
			StudentType:        parseInt(row[11]),
			CheckInTime:        row[12],
			RegistrationTime:   row[13],
			Academic:           parseInt(row[14]),
		}

		// 添加或更新学生信息
		if err := utils.DB_MySQL.Save(&student).Error; err != nil {
			return fmt.Errorf("保存学生信息失败: %w", err)
		}
	}
	return nil
}

func importTeachers(f *excelize.File) error {
	rows, err := f.GetRows("Teachers")
	if err != nil {
		return err
	}

	for _, row := range rows[1:] {
		teacher := models.UserAccount{
			Account:  row[0],
			Password: row[1],
			ChatType: parseInt(row[2]),
		}

		// 添加或更新教师账号
		if err := utils.DB_MySQL.Save(&teacher).Error; err != nil {
			return fmt.Errorf("保存教师账号信息失败: %w", err)
		}

		contact := models.ContactInformation{
			Account:               row[0],
			CorrespondenceAddress: row[3],
			Phone:                 row[4],
			Email:                 row[5],
			Landline:              row[6],
			PostCode:              row[7],
			HomeAddress:           row[8],
		}

		// 添加或更新教师联系方式
		if err := utils.DB_MySQL.Save(&contact).Error; err != nil {
			return fmt.Errorf("保存教师联系方式失败: %w", err)
		}
	}
	return nil
}

func importAdmins(f *excelize.File) error {
	rows, err := f.GetRows("Admins")
	if err != nil {
		return err
	}

	for _, row := range rows[1:] {
		admin := models.UserAccount{
			Account:  row[0],
			Password: row[1],
			ChatType: parseInt(row[2]),
		}

		// 添加或更新管理员账号
		if err := utils.DB_MySQL.Save(&admin).Error; err != nil {
			return fmt.Errorf("保存管理员账号信息失败: %w", err)
		}

		contact := models.ContactInformation{
			Account:               row[0],
			CorrespondenceAddress: row[3],
			Phone:                 row[4],
			Email:                 row[5],
			Landline:              row[6],
			PostCode:              row[7],
			HomeAddress:           row[8],
		}

		// 添加或更新管理员联系方式
		if err := utils.DB_MySQL.Save(&contact).Error; err != nil {
			return fmt.Errorf("保存管理员联系方式失败: %w", err)
		}
	}
	return nil
}

// 工具函数
func parseBool(s string) int {
	if s == "true" {
		return 1
	}
	return 0
}

func parseInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return value
}
