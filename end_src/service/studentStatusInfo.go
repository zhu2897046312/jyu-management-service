package service

import (
	"jyu-service/models"
	"jyu-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ContactsResponse 定义
// StudentStatusInfoResponse 定义
type StudentStatusInfoResponse struct {
	models.StudentStatusInformation
	UserName string `json:"user_name"` // 姓名 (从 UserBasicInformation 表中获取)
}

// GetStudentStatusInfo 获取学籍信息列表（支持分页）
func GetStudentStatusInfo(c *gin.Context) {
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

	var studentStatusInfos []StudentStatusInfoResponse
	var total int64

	// 获取总数
	if err := utils.DB_MySQL.Model(&models.StudentStatusInformation{}).Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "无法获取学籍信息总数",
			"error":   err.Error(),
		})
		return
	}

	// 联表查询 StudentStatusInformation 和 UserBasicInformation，获取 UserName
	if err := utils.DB_MySQL.Table("student_status_informations").
		Select("student_status_informations.*, user_basic_informations.name as user_name").
		Joins("left join user_basic_informations on student_status_informations.account = user_basic_informations.account").
		Limit(pageSize).
		Offset(offset).
		Scan(&studentStatusInfos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "无法获取学籍信息数据",
			"error":   err.Error(),
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"status":              "success",
		"StudentStatusInfo":    studentStatusInfos,
		"total":               total,
	})
}


// AddCourse 添加学籍信息信息
func AddStudentStatusInfo(c *gin.Context) {
    var courseInfo models.StudentStatusInformation

    if err := c.ShouldBindJSON(&courseInfo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
			"status": "error", 
			"message": "请求数据格式错误", 
			"error": err.Error(),
		})
        return
    }

    var existingCourse models.StudentStatusInformation
    if err := utils.DB_MySQL.Where("account = ?", courseInfo.Account).First(&existingCourse).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{
			"status": "error", 
			"message": "学籍信息代码已存在",
		})
        return
    }

    if err := utils.DB_MySQL.Create(&courseInfo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "添加学籍信息失败", 
			"error": err.Error(),
		})
        return
    }

    c.JSON(http.StatusOK, gin.H{
		"status": "success", 
		"message": "学籍信息添加成功", 
		"course": courseInfo,
	})
}

// UpdateUserBasic 更新学籍信息信息
func UpdateStudentStatusInfo(c *gin.Context) {
    var courseInfo models.StudentStatusInformation
    if err := c.ShouldBindJSON(&courseInfo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
			"status": "error", 
			"message": "请求数据格式错误", 
			"error": err.Error(),
		})
        return
    }

    if err := utils.DB_MySQL.Model(&models.StudentStatusInformation{}).Where("account = ?", courseInfo.Account).Updates(courseInfo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "更新学籍信息失败", 
			"error": err.Error(),
		})
        return
    }

    c.JSON(http.StatusOK, gin.H{
		"status": "success", 
		"message": "学籍信息信息更新成功", 
		"course": courseInfo,
	})
}

// DeleteUserBasic 删除学籍信息信息
func DeleteStudentStatusInfo(c *gin.Context) {
	courseCode := c.Param("account")

	var course models.StudentStatusInformation
	if err := utils.DB_MySQL.Where("account = ?", courseCode).First(&course).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error", 
			"message": "未找到该学籍信息", 
			"error": err.Error(),
		})
		return
	}

	if err := utils.DB_MySQL.Delete(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "删除学籍信息失败", 
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success", 
		"message": "学籍信息删除成功", 
		"course": course,
	})
}
