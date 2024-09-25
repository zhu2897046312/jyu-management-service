package service

import (
	"fmt"
	"jyu-service/models"
	"jyu-service/utils"
	"log"
	"net/http"
	"strconv"
	"gorm.io/gorm"
	// "os"
	// "io"
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

// 生成 Excel 列名 (A, B, ..., Z, AA, AB, ..., ZZ, AAA, etc.)
func getExcelColumnName(n int) string {
    colName := ""
    for n > 0 {
        n-- // 将 n 转换为从 0 开始
        colName = string(rune('A'+(n%26))) + colName
        n /= 26
    }
    return colName
}

// 生成固定格式的 Excel 文件
func GenerateExcelTemplate(c *gin.Context) {
	f := excelize.NewFile()

	// 创建“Students” Sheet 并写入表头
	index, err := f.NewSheet("Sheet1")
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
		column := getExcelColumnName(i + 1) // i 从 0 开始，所以要加 1
		f.SetCellValue("Sheet1", column+"1", header)
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

// 解析上传的 Excel 文件并根据 account 存在与否更新或插入信息
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

	// 3. 使用 excelize.OpenReader 读取文件内容
	excelFile, err := excelize.OpenReader(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法解析 Excel 文件"})
		return
	}

	// 确保获取并使用正确的工作表名 "Sheet1"
	sheetName := "Sheet1"
	rows, err := excelFile.GetRows(sheetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法读取工作表"})
		return
	}

	// 4. 遍历 Excel 文件中的每一行，从第二行开始
	for i, row := range rows {
		if i == 0 {
			// 跳过表头
			continue
		}

		// 获取账号
		account := row[0]
		chatType, _ := strconv.Atoi(row[2])

		// 5. 根据 account 检查用户是否已存在
		var userAccount models.UserAccount
		result := utils.DB_MySQL.Where("account = ?", account).First(&userAccount)

		if result.RowsAffected > 0 {
			// 账号已存在，执行更新操作
			userAccount.Password = row[1]
			userAccount.ChatType = chatType
			utils.DB_MySQL.Save(&userAccount)
		} else {
			// 账号不存在，执行插入操作
			userAccount = models.UserAccount{
				Password: row[1],
				ChatType: chatType,
			}
			account_db , db := userAccount.Insert_auto()
			if db.Error!= nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
                return
			}
			account = account_db
		}

		// 插入或更新基本信息
		name := row[3]
		sex, _ := strconv.Atoi(row[4])
		identificationType := row[5]
		identificationNumber := row[6]
		ethnicGroup := row[7]
		birthday := row[8]
		oldName := row[9]
		politicalOutlook := row[10]
		enrollmentDates := row[11]

		var basicInfo models.UserBasicInformation
		result = utils.DB_MySQL.Where("account = ?", account).First(&basicInfo)

		if result.RowsAffected > 0 {
			// 更新基本信息
			basicInfo.Name = name
			basicInfo.Sex = sex
			basicInfo.IdentificationType = identificationType
			basicInfo.IdentificationNumber = identificationNumber
			basicInfo.EthnicGroup = ethnicGroup
			basicInfo.Birthday = birthday
			basicInfo.OldName = oldName
			basicInfo.PoliticalOutlook = politicalOutlook
			basicInfo.EnrollmentDates = enrollmentDates
			utils.DB_MySQL.Save(&basicInfo)
		} else {
			// 插入基本信息
			basicInfo = models.UserBasicInformation{
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
		}

		// 插入或更新联系方式
		correspondenceAddress := row[12]
		phone := row[13]
		email := row[14]
		landline := row[15]
		postCode := row[16]
		homeAddress := row[17]

		var contactInfo models.ContactInformation
		result = utils.DB_MySQL.Where("account = ?", account).First(&contactInfo)

		if result.RowsAffected > 0 {
			// 更新联系方式
			contactInfo.CorrespondenceAddress = correspondenceAddress
			contactInfo.Phone = phone
			contactInfo.Email = email
			contactInfo.Landline = landline
			contactInfo.PostCode = postCode
			contactInfo.HomeAddress = homeAddress
			utils.DB_MySQL.Save(&contactInfo)
		} else {
			// 插入联系方式
			contactInfo = models.ContactInformation{
				Account:               account,
				CorrespondenceAddress: correspondenceAddress,
				Phone:                 phone,
				Email:                 email,
				Landline:              landline,
				PostCode:              postCode,
				HomeAddress:           homeAddress,
			}
			utils.DB_MySQL.Create(&contactInfo)
		}

		// 如果 chatType 为学生，更新或插入学籍信息
		if chatType == models.Student {
			academicYear := row[18]
			academyName := row[19]
			className := row[20]
			professionalName := row[21]
			status := row[22]
			isInSchool, _ := strconv.Atoi(row[23])
			registrationStatus := row[24]
			educationalLevel := row[25]
			cultivationMethod := row[26]
			cultivationLevel, _ := strconv.Atoi(row[27])
			studentType, _ := strconv.Atoi(row[28])
			checkInTime := row[29]
			registrationTime := row[30]
			academic, _ := strconv.Atoi(row[31])

			var studentStatus models.StudentStatusInformation
			result = utils.DB_MySQL.Where("account = ?", account).First(&studentStatus)

			if result.RowsAffected > 0 {
				// 更新学籍信息
				studentStatus.AcademicYear = academicYear
				studentStatus.AcademyName = academyName
				studentStatus.ClassName = className
				studentStatus.ProfessionalName = professionalName
				studentStatus.Status = status
				studentStatus.IsInSchool = isInSchool
				studentStatus.RegistrationStatus = registrationStatus
				studentStatus.EducationalLevel = educationalLevel
				studentStatus.CultivationMethod = cultivationMethod
				studentStatus.CultivationLevel = cultivationLevel
				studentStatus.StudentType = studentType
				studentStatus.CheckInTime = checkInTime
				studentStatus.RegistrationTime = registrationTime
				studentStatus.Academic = academic
				utils.DB_MySQL.Save(&studentStatus)
			} else {
				// 插入学籍信息
				studentStatus = models.StudentStatusInformation{
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
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Excel 解析并更新/生成账号成功",
	})
}

// 获取所有信息并生成 Excel 文件
func GetAllInfoExecl(c *gin.Context) {
	var accounts []models.UserAccount

	// 1. 查询所有账号
	if err := utils.DB_MySQL.Find(&accounts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取账号失败"})
		return
	}

	// 2. 创建 Excel 文件
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.NewSheet(sheet)

	// 设置表头
	headers := []string{
		"账号", "密码", "账户类型", "姓名", "性别", "身份证号", "出生日期",
		"民族", "证件类型", "曾用名", "政治面貌", "入学日期",  "通讯地址",
		"手机号码", "电子邮箱", "固定电话", "邮政编码", "家庭地址",
		"年级", "学院名称",
		"班级名称", "专业名称", "学籍状态", "是否在校", "报到注册状态", "学历层次",
		"培养方式", "培养层次", "学生类别", "报到时间", "注册时间", "学制",
	}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, header)
	}

	// 3. 查询每个账号的详细信息并填充到 Excel 中
	for idx, account := range accounts {
		row := idx + 2

		// 填入账号数据
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), account.Account)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), account.Password)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), account.ChatType)

		// 基本信息
		var basicInfo models.UserBasicInformation
		if err := utils.DB_MySQL.Where("account = ?", account.Account).First(&basicInfo).Error; err == nil {
			// 填入基本信息
			f.SetCellValue(sheet, fmt.Sprintf("D%d", row), basicInfo.Name)
			f.SetCellValue(sheet, fmt.Sprintf("E%d", row), basicInfo.Sex)
			f.SetCellValue(sheet, fmt.Sprintf("F%d", row), basicInfo.IdentificationNumber)
			f.SetCellValue(sheet, fmt.Sprintf("G%d", row), basicInfo.Birthday)
			f.SetCellValue(sheet, fmt.Sprintf("H%d", row), basicInfo.EthnicGroup)
			f.SetCellValue(sheet, fmt.Sprintf("I%d", row), basicInfo.IdentificationType)
			f.SetCellValue(sheet, fmt.Sprintf("J%d", row), basicInfo.OldName)
			f.SetCellValue(sheet, fmt.Sprintf("K%d", row), basicInfo.PoliticalOutlook)
			f.SetCellValue(sheet, fmt.Sprintf("L%d", row), basicInfo.EnrollmentDates)
		}

		// 联系方式
		var contactInfo models.ContactInformation
		if err := utils.DB_MySQL.Where("account = ?", account.Account).First(&contactInfo).Error; err == nil {
			// 填入联系方式
			f.SetCellValue(sheet, fmt.Sprintf("M%d", row), contactInfo.CorrespondenceAddress)
			f.SetCellValue(sheet, fmt.Sprintf("N%d", row), contactInfo.Phone)
			f.SetCellValue(sheet, fmt.Sprintf("O%d", row), contactInfo.Email)
			f.SetCellValue(sheet, fmt.Sprintf("P%d", row), contactInfo.Landline)
			f.SetCellValue(sheet, fmt.Sprintf("Q%d", row), contactInfo.PostCode)
			f.SetCellValue(sheet, fmt.Sprintf("R%d", row), contactInfo.HomeAddress)
		}

		// 学籍信息
		var statusInfo models.StudentStatusInformation
		if err := utils.DB_MySQL.Where("account = ?", account.Account).First(&statusInfo).Error; err == nil {
			// 填入学籍信息
			f.SetCellValue(sheet, fmt.Sprintf("S%d", row), statusInfo.AcademicYear)
			f.SetCellValue(sheet, fmt.Sprintf("T%d", row), statusInfo.AcademyName)
			f.SetCellValue(sheet, fmt.Sprintf("U%d", row), statusInfo.ClassName)
			f.SetCellValue(sheet, fmt.Sprintf("V%d", row), statusInfo.ProfessionalName)
			f.SetCellValue(sheet, fmt.Sprintf("W%d", row), statusInfo.Status)
			f.SetCellValue(sheet, fmt.Sprintf("X%d", row), statusInfo.IsInSchool)
			f.SetCellValue(sheet, fmt.Sprintf("Y%d", row), statusInfo.RegistrationStatus)
			f.SetCellValue(sheet, fmt.Sprintf("Z%d", row), statusInfo.EducationalLevel)
			f.SetCellValue(sheet, fmt.Sprintf("AA%d", row), statusInfo.CultivationMethod)
			f.SetCellValue(sheet, fmt.Sprintf("AB%d", row), statusInfo.CultivationLevel)
			f.SetCellValue(sheet, fmt.Sprintf("AC%d", row), statusInfo.StudentType)
			f.SetCellValue(sheet, fmt.Sprintf("AD%d", row), statusInfo.CheckInTime)
			f.SetCellValue(sheet, fmt.Sprintf("AE%d", row), statusInfo.RegistrationTime)
			f.SetCellValue(sheet, fmt.Sprintf("AF%d", row), statusInfo.Academic)
		}
	}

	// 4. 生成文件并返回给前端
	// 设置 Excel 输出流
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=all.xlsx")
	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// execl 生成课程信息文档
func GenerateCourseExcel(c *gin.Context) {
	f := excelize.NewFile()

	// 创建“Students” Sheet 并写入表头
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		log.Fatal(err)
	}

	// 公共表头：账号信息、基本信息、联系方式
	commonHeaders := []string{
		"课程代码","课程名称","教师账号","教师名称",
		"上课时间（节次:星期:持续周次）","上课地点",
		"学分","教学模式","年级","学期",
		"开课学院","课程归属","课程类别","课程性质",
		"教学班名称","总人数","已选人数",
	}

	// 写入学生表头
	for i, header := range commonHeaders {
		column := getExcelColumnName(i + 1) // i 从 0 开始，所以要加 1
		f.SetCellValue("Sheet1", column+"1", header)
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

func GetCoursesInfoExecl(c *gin.Context) {
	var accounts []models.CourseInformation

	// 1. 查询所有账号
	if err := utils.DB_MySQL.Find(&accounts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取账号失败"})
		return
	}

	// 2. 创建 Excel 文件
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.NewSheet(sheet)

	// 设置表头
	headers := []string{
		"课程代码","课程名称","教师账号","教师名称",
		"上课时间（节次:星期:持续周次）","上课地点",
		"学分","教学模式","年级","学期",
		"开课学院","课程归属","课程类别","课程性质",
		"教学班名称","总人数","已选人数",
	}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, header)
	}

	// 3. 查询每个账号的详细信息并填充到 Excel 中
	for idx, account := range accounts {
		row := idx + 2

		// 填入账号数据
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), account.CourseCode)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), account.CourseName)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), account.Account)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), account.TeacherName)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), account.ClassTime)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), account.ClassAddress)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), account.Credits)
		f.SetCellValue(sheet, fmt.Sprintf("H%d", row), account.TeachingMode)
		f.SetCellValue(sheet, fmt.Sprintf("I%d", row), account.AcademicYear)
		f.SetCellValue(sheet, fmt.Sprintf("J%d", row), account.Semester)
		f.SetCellValue(sheet, fmt.Sprintf("K%d", row), account.CommencementAcademy)
		f.SetCellValue(sheet, fmt.Sprintf("L%d", row), account.CourseAffiliation)
		f.SetCellValue(sheet, fmt.Sprintf("M%d", row), account.CourseType)
		f.SetCellValue(sheet, fmt.Sprintf("N%d", row), account.CourseNature)
		f.SetCellValue(sheet, fmt.Sprintf("O%d", row), account.ClassName)
		f.SetCellValue(sheet, fmt.Sprintf("P%d", row), account.MaxStudentNumber)
		f.SetCellValue(sheet, fmt.Sprintf("Q%d", row), account.ChoosedNumber)
	}

	// 4. 生成文件并返回给前端
	// 设置 Excel 输出流
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=all.xlsx")
	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// 解析 Excel 文件并处理课程信息
func ParseCourseExcel(c *gin.Context) {
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

	// 3. 使用 excelize.OpenReader 读取文件内容
	excelFile, err := excelize.OpenReader(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法解析 Excel 文件"})
		return
	}

	// 获取所有行数据
	rows, err := excelFile.GetRows("Sheet1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法读取 Excel 行数据"})
		return
	}

	// 遍历行数据，更新或插入课程信息
	for _, row := range rows[1:] { // 跳过表头
		if len(row) < 15 { // 确保行数据完整
			continue
		}

		course := models.CourseInformation{
			CourseCode:          row[0],
			CourseName:          row[1],
			Account:             row[2],
			TeacherName:         row[3],
			ClassTime:           row[4],
			ClassAddress:        row[5],
			Credits:             parseFloat(row[6]),
			TeachingMode:        parseInt(row[7]),
			AcademicYear:        row[8],
			Semester:            parseInt(row[9]),
			CommencementAcademy: row[10],
			CourseAffiliation:   row[11],
			CourseType:          parseInt(row[12]),
			CourseNature:        parseInt(row[13]),
			ClassName:           row[14],
			MaxStudentNumber:    parseInt(row[15]),
		}

		// 检查数据库中是否已存在该课程
		var existingCourse models.CourseInformation
		if err := utils.DB_MySQL.Where("course_code = ?", course.CourseCode).First(&existingCourse).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// 课程不存在，插入新记录
				if err := utils.DB_MySQL.Create(&course).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("插入课程 %s 失败: %v", course.CourseCode, err)})
					return
				}
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查询失败"})
				return
			}
		} else {
			// 课程已存在，更新记录
			if err := utils.DB_MySQL.Model(&existingCourse).Updates(course).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("更新课程 %s 失败: %v", course.CourseCode, err)})
				return
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "课程信息处理完成"})
}

// 辅助函数：将字符串转换为整数
func parseInt(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}

// 辅助函数：将字符串转换为浮点数
func parseFloat(s string) float32 {
	value, _ := strconv.ParseFloat(s, 32)
	return float32(value)
}

func GenerateStudentInfoExcel(c *gin.Context) {
	courseCode := c.Query("course_code")

	var userCourses []models.UserCourse
    // 根据 course_code 查询 UserCourse
    if err := utils.DB_MySQL.Where("course_code = ?", courseCode).Find(&userCourses).Error; err != nil {
        return 
    }

    // 创建 Excel 文件
    f := excelize.NewFile()
    sheet := "Sheet1"
    f.NewSheet(sheet)

    // 设置表头
	headers := []string{
		"课程代码","学号","专业名称","班级名称","成绩",
	}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, header)
	}

	// 3. 查询每个账号的详细信息并填充到 Excel 中
	for idx, account := range userCourses {
		row := idx + 2

		// 查询学生个人信息
		var student models.StudentStatusInformation
		if err := utils.DB_MySQL.Where("account = ?", account.Account).Find(&student).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询学生信息失败"})
			return
		}
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), account.CourseCode)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), account.Account)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), student.ProfessionalName)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), student.ClassName)

		if account.CourseGrade != "" {
			f.SetCellValue(sheet, fmt.Sprintf("E%d", row), account.CourseGrade)
		}
	}

	// 4. 生成文件并返回给前端
	// 设置 Excel 输出流
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=all.xlsx")
	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// 解析 Excel 文件 修改成绩
func UploadStudentScores(c *gin.Context) {
    // 获取上传的文件
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "文件获取失败"})
        return
    }

    // 打开上传的 Excel 文件
    excelFile, err := file.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "打开文件失败"})
        return
    }
    defer excelFile.Close()

    // 读取 Excel 文件
    f, err := excelize.OpenReader(excelFile)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "解析 Excel 文件失败"})
        return
    }

    // 获取第一个工作表的名称
    sheetName := f.GetSheetName(0)

    // 读取表格数据
    rows, err := f.GetRows(sheetName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "读取行数据失败"})
        return
    }

    // 跳过表头，从第二行开始
    for i, row := range rows {
        if i == 0 {
            continue // 跳过表头
        }

        courseCode := row[0]      // 课程代码
        account := row[1]         // 学号
        courseGrade := row[4]     // 成绩

        // 更新数据库
        var userCourse models.UserCourse
        if err := utils.DB_MySQL.Where("account = ? AND course_code = ?", account, courseCode).First(&userCourse).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "更新成绩失败: 学生信息不存在"})
            return
        }

        // 更新成绩
        userCourse.CourseGrade = courseGrade
        if err := utils.DB_MySQL.Save(&userCourse).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "更新成绩失败"})
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{"message": "成绩上传成功"})
}




