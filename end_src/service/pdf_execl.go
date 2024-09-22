package service

import (
	"fmt"
	"jyu-service/models"
	"jyu-service/utils"
	"log"
	"net/http"
	"strconv"
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
		column := getExcelColumnName(i + 1) // i 从 0 开始，所以要加 1
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
	fmt.Println(" index:" + string(index))
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
			fmt.Println(err.Error())
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

	c.JSON(
		http.StatusOK, gin.H{
            "message": "Excel 解析并生成账号成功",
        },
	)
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
		"民族", "证件类型", "曾用名", "政治面貌", "入学日期", "年级", "学院名称",
		"班级名称", "专业名称", "学籍状态", "是否在校", "报到注册状态", "学历层次",
		"培养方式", "培养层次", "学生类别", "报到时间", "注册时间", "学制", "通讯地址",
		"手机号码", "电子邮箱", "固定电话", "邮政编码", "家庭地址",
	}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, header)
	}

	// 3. 查询每个账号的详细信息并填充到 Excel 中
	for idx, account := range accounts {
		// 基本信息
		var basicInfo models.UserBasicInformation
		utils.DB_MySQL.Where("account = ?", account.Account).First(&basicInfo)

		// 学籍信息
		var statusInfo models.StudentStatusInformation
		utils.DB_MySQL.Where("account = ?", account.Account).First(&statusInfo)

		// 联系方式
		var contactInfo models.ContactInformation
		utils.DB_MySQL.Where("account = ?", account.Account).First(&contactInfo)

		// 填入数据
		row := idx + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), account.Account)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), account.Password)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), account.ChatType)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), basicInfo.Name)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), basicInfo.Sex)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), basicInfo.IdentificationNumber)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), basicInfo.Birthday)
		f.SetCellValue(sheet, fmt.Sprintf("H%d", row), basicInfo.EthnicGroup)
		f.SetCellValue(sheet, fmt.Sprintf("I%d", row), basicInfo.IdentificationType)
		f.SetCellValue(sheet, fmt.Sprintf("J%d", row), basicInfo.OldName)
		f.SetCellValue(sheet, fmt.Sprintf("K%d", row), basicInfo.PoliticalOutlook)
		f.SetCellValue(sheet, fmt.Sprintf("L%d", row), basicInfo.EnrollmentDates)

		// 学籍信息
		f.SetCellValue(sheet, fmt.Sprintf("M%d", row), statusInfo.AcademicYear)
		f.SetCellValue(sheet, fmt.Sprintf("N%d", row), statusInfo.AcademyName)
		f.SetCellValue(sheet, fmt.Sprintf("O%d", row), statusInfo.ClassName)
		f.SetCellValue(sheet, fmt.Sprintf("P%d", row), statusInfo.ProfessionalName)
		f.SetCellValue(sheet, fmt.Sprintf("Q%d", row), statusInfo.Status)
		f.SetCellValue(sheet, fmt.Sprintf("R%d", row), statusInfo.IsInSchool)
		f.SetCellValue(sheet, fmt.Sprintf("S%d", row), statusInfo.RegistrationStatus)
		f.SetCellValue(sheet, fmt.Sprintf("T%d", row), statusInfo.EducationalLevel)
		f.SetCellValue(sheet, fmt.Sprintf("U%d", row), statusInfo.CultivationMethod)
		f.SetCellValue(sheet, fmt.Sprintf("V%d", row), statusInfo.CultivationLevel)
		f.SetCellValue(sheet, fmt.Sprintf("W%d", row), statusInfo.StudentType)
		f.SetCellValue(sheet, fmt.Sprintf("X%d", row), statusInfo.CheckInTime)
		f.SetCellValue(sheet, fmt.Sprintf("Y%d", row), statusInfo.RegistrationTime)
		f.SetCellValue(sheet, fmt.Sprintf("Z%d", row), statusInfo.Academic)

		// 联系方式
		f.SetCellValue(sheet, fmt.Sprintf("AA%d", row), contactInfo.CorrespondenceAddress)
		f.SetCellValue(sheet, fmt.Sprintf("AB%d", row), contactInfo.Phone)
		f.SetCellValue(sheet, fmt.Sprintf("AC%d", row), contactInfo.Email)
		f.SetCellValue(sheet, fmt.Sprintf("AD%d", row), contactInfo.Landline)
		f.SetCellValue(sheet, fmt.Sprintf("AE%d", row), contactInfo.PostCode)
		f.SetCellValue(sheet, fmt.Sprintf("AF%d", row), contactInfo.HomeAddress)
	}

	// 4. 生成文件并返回给前端
	// 设置 Excel 输出流
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=all.xlsx")
	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}




