package service

import (
	"jyu-service/models"
	"jyu-service/utils"
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
