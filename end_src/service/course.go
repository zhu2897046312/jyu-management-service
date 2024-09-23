package service

import (
	"jyu-service/models"
	"jyu-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"gorm.io/gorm"
)
/**
	admin
*/

// GetCourses 获取课程列表（支持分页）
func GetCourses(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	var courses []models.CourseInformation
	var total int64

	if err := utils.DB_MySQL.Model(&models.CourseInformation{}).Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "无法获取课程总数", 
			"error": err.Error(),
		})
		return
	}

	if err := utils.DB_MySQL.Limit(pageSize).Offset(offset).Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "无法获取课程数据", 
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"courses": courses,
		"total":   total,
	})
}

// AddCourse 添加课程信息
func AddCourse(c *gin.Context) {
    var courseInfo models.CourseInformation

    if err := c.ShouldBindJSON(&courseInfo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
			"status": "error", 
			"message": "请求数据格式错误", 
			"error": err.Error(),
		})
        return
    }

    var existingCourse models.CourseInformation
    if err := utils.DB_MySQL.Where("course_code = ?", courseInfo.CourseCode).First(&existingCourse).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{
			"status": "error", 
			"message": "课程代码已存在",
		})
        return
    }

    if err := utils.DB_MySQL.Create(&courseInfo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "添加课程失败", 
			"error": err.Error(),
		})
        return
    }

    c.JSON(http.StatusOK, gin.H{
		"status": "success", 
		"message": "课程添加成功", 
		"course": courseInfo,
	})
}

// UpdateCourses 更新课程信息
func UpdateCourses(c *gin.Context) {
    var courseInfo models.CourseInformation
    if err := c.ShouldBindJSON(&courseInfo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
			"status": "error", 
			"message": "请求数据格式错误", 
			"error": err.Error(),
		})
        return
    }

    if err := utils.DB_MySQL.Model(&models.CourseInformation{}).Where("course_code = ?", courseInfo.CourseCode).Updates(courseInfo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "更新课程失败", 
			"error": err.Error(),
		})
        return
    }

    c.JSON(http.StatusOK, gin.H{
		"status": "success", 
		"message": "课程信息更新成功", 
		"course": courseInfo,
	})
}

// DeleteCourses 删除课程信息
func DeleteCourses(c *gin.Context) {
	courseCode := c.Param("course_code")

	var course models.CourseInformation
	if err := utils.DB_MySQL.Where("course_code = ?", courseCode).First(&course).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error", 
			"message": "未找到该课程", 
			"error": err.Error(),
		})
		return
	}

	if err := utils.DB_MySQL.Delete(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "删除课程失败", 
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success", 
		"message": "课程删除成功", 
		"course": course,
	})
}

/**
	teacher -> course
*/

func GetTeacherCourse(c *gin.Context) {
	account_teacher := c.Query("account")

	// 根据教师名称查找课程
    var courses []models.CourseInformation
    if err := utils.DB_MySQL.Where("accountaccount = ?", account_teacher).Find(&courses).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status": "error", 
            "message": "查询课程失败", 
            "error": err.Error(),
        })
        return
    }

	// 成功返回课程列表
    c.JSON(http.StatusOK, gin.H{
        "status": "success", 
        "Courses": courses,
    })
}

func ModifinedGrades(c *gin.Context) {
	var userCourse models.UserCourse

    // 解析请求体中的 JSON 数据
    if err := c.ShouldBindJSON(&userCourse); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status": "error",
            "message": "请求数据格式错误",
            "error": err.Error(),
        })
        return
    }

    // 确保 account 和 course_code 存在
    if userCourse.Account == "" || userCourse.CourseCode == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "status": "error",
            "message": "学号和课程代码不能为空",
        })
        return
    }

    // 查找该学生是否已经选了该课程
    var existingRecord models.UserCourse
    if err := utils.DB_MySQL.Where("account = ? AND course_code = ?", userCourse.Account, userCourse.CourseCode).First(&existingRecord).Error; err != nil {
        // 如果找不到记录，返回错误信息
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{
                "status": "error",
                "message": "未找到该学生的课程记录",
            })
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{
                "status": "error",
                "message": "查询学生课程记录失败",
                "error": err.Error(),
            })
        }
        return
    }

    // 更新成绩信息
    existingRecord.CourseGrade = userCourse.CourseGrade
    if err := utils.DB_MySQL.Save(&existingRecord).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status": "error",
            "message": "更新成绩失败",
            "error": err.Error(),
        })
        return
    }

    // 成功返回结果
    c.JSON(http.StatusOK, gin.H{
        "status": "success",
        "message": "成绩更新成功",
        "data": existingRecord,
    })
}

  