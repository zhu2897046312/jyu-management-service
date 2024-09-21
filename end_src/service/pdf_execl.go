package service

import (
	"fmt"
	"jyu-service/models"
	"jyu-service/utils"
	"log"
	"net/http"
	"strconv"
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
				f.SetCellValue(sheetName, "G"+strconv.Itoa(row), CalculateGradePoints(uc.CourseGrade, c.Credits))
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
	index, err := f.NewSheet("Template")
	if err != nil {
		log.Fatal(err)
	}

	// 公共表头：账号信息、基本信息、联系方式
	commonHeaders := []string{
		"账号", "密码", "账号类型", // 账号信息
		"姓名", "性别", "身份证件类型", "身份证件号", // 基本信息
		"民族", "出生日期", "曾用名", // 基本信息
		"政治面貌", "入学日期", // 基本信息
		"通讯地址", "手机号码", "邮箱", // 联系方式
		"固定电话", "邮政编码", "家庭住址", // 联系方式
	}

	// 学生特有字段
	studentSpecificHeaders := []string{
		"年级", "学院名称", "班级名称", // 学籍信息
		"专业名称", "学籍状态", "是否在校", // 学籍信息
		"报到注册状态", "学历层次", // 学籍信息
		"培养方式", "培养层次", // 学籍信息
		"学生类别", "报到时间", "注册时间", // 学籍信息
		"学制", // 学籍信息
	}

	// 组合学生表头
	studentHeaders := append(commonHeaders, studentSpecificHeaders...)

	// 写入学生表头
	for i, header := range studentHeaders {
		column := string('A' + i)
		f.SetCellValue("Template", column+"1", header)
	}

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

// 解析上传的 Excel 文件并生成账号
// 解析上传的 Excel 文件并生成账号
func ImportAndGenerateAccounts(c *gin.Context) {
	// 1. 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无法获取文件"})
		return
	}

	// 2. 打开上传的文件
	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法打开文件"})
		return
	}
	defer f.Close()

	// 4. 使用 excelize.OpenReader 读取文件内容
	excelFile, err := excelize.OpenReader(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法解析 Excel 文件"})
		return
	}

	// 确保获取并使用正确的工作表名 "Template"
	sheetName := "Template"
	index, err := excelFile.GetSheetIndex(sheetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取工作表索引"})
		return
	}
	if index == -1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "工作表 Template 不存在"})
		return
	}

	// 4. 获取 Students Sheet 的内容
	rows, err := excelFile.GetRows(sheetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法读取工作表"})
		return
	}

	// 5. 遍历 Excel 文件中的每一行，从第二行开始
	for i, row := range rows {
		if i == 0 {
			// 跳过表头
			continue
		}
		chatType, _ := strconv.Atoi(row[2])
		
		// 将数据插入数据库
		userAccount := models.UserAccount{
			Password: row[1],
			ChatType: chatType,
		}

		// 插入到 UserAccount 表中
		account, db := userAccount.Insert_auto()
		if db.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
			return
		}

		// 在 Excel 中写入生成的账号到指定的单元格，比如 A 列
		cell := fmt.Sprintf("A%d", i+1)
		fmt.Printf("Writing account %s to cell %s\n", account, cell)
		err := excelFile.SetCellValue(sheetName, cell, account)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}

		// 处理其他信息并插入相应的表
		name := row[3]
		sex, _ := strconv.Atoi(row[4])
		identificationType := row[5]
		identificationNumber := row[6]
		ethnicGroup := row[7]
		birthday := row[8]
		oldName := row[9]
		politicalOutlook := row[10]
		enrollmentDates := row[11]

		// 插入 UserBasicInformation
		basicInfo := models.UserBasicInformation{
			Account:              account,
			Name:                 name,
			Sex:                  sex,
			IdentificationType:   identificationType,
			IdentificationNumber: identificationNumber,
			EthnicGroup:          ethnicGroup,
			Birthday:             birthday,
			OldName:              oldName,
			PoliticalOutlook:     politicalOutlook,
			EnrollmentDates:      enrollmentDates,
		}
		utils.DB_MySQL.Create(&basicInfo)

		// 插入联系方式
		correspondenceAddress := row[12]
		phone := row[13]
		email := row[14]
		landline := row[15]
		postCode := row[16]
		homeAddress := row[17]

		contactInfo := models.ContactInformation{
			Account:               account,
			CorrespondenceAddress: correspondenceAddress,
			Phone:                 phone,
			Email:                 email,
			Landline:              landline,
			PostCode:              postCode,
			HomeAddress:           homeAddress,
		}
		utils.DB_MySQL.Create(&contactInfo)

		if chatType == models.Student {
			academicYear := row[12]
			academyName := row[13]
			className := row[14]
			professionalName := row[15]
			status := row[16]
			isInSchool, _ := strconv.Atoi(row[17])
			registrationStatus := row[18]
			educationalLevel := row[19]
			cultivationMethod := row[20]
			cultivationLevel, _ := strconv.Atoi(row[21])
			studentType, _ := strconv.Atoi(row[22])
			checkInTime := row[23]
			registrationTime := row[24]
			academic, _ := strconv.Atoi(row[25])

			studentStatus := models.StudentStatusInformation{
				Account:            account,
				AcademicYear:       academicYear,
				AcademyName:        academyName,
				ClassName:          className,
				ProfessionalName:   professionalName,
				Status:             status,
				IsInSchool:         isInSchool,
				RegistrationStatus: registrationStatus,
				EducationalLevel:   educationalLevel,
				CultivationMethod:  cultivationMethod,
				CultivationLevel:   cultivationLevel,
				StudentType:        studentType,
				CheckInTime:        checkInTime,
				RegistrationTime:   registrationTime,
				Academic:           academic,
			}
			utils.DB_MySQL.Create(&studentStatus)
		}
	}

	// 设置响应头，返回文件流给前端
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=\"update_template.xlsx\"")
	c.Header("File-Name", "update_template.xlsx")
	if err := excelFile.Write(c.Writer); err != nil {
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
