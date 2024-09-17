package service

import (
	"jyu-service/models"
	"jyu-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUserBasic 获取用户基本信息列表（支持分页）
func GetUserBasicInfo(c *gin.Context) {
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

	var UserBasic []models.UserBasicInformation
	var total int64

	if err := utils.DB_MySQL.Model(&models.UserBasicInformation{}).Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "无法获取用户基本信息总数", 
			"error": err.Error(),
		})
		return
	}

	if err := utils.DB_MySQL.Limit(pageSize).Offset(offset).Find(&UserBasic).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "无法获取用户基本信息数据", 
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"UserBasic": UserBasic,
		"total":   total,
	})
}

// AddCourse 添加用户基本信息信息
func AddUserBasicInfo(c *gin.Context) {
    var courseInfo models.UserBasicInformation

    if err := c.ShouldBindJSON(&courseInfo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
			"status": "error", 
			"message": "请求数据格式错误", 
			"error": err.Error(),
		})
        return
    }

    var existingCourse models.UserBasicInformation
    if err := utils.DB_MySQL.Where("account = ?", courseInfo.Account).First(&existingCourse).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{
			"status": "error", 
			"message": "用户基本信息代码已存在",
		})
        return
    }

    if err := utils.DB_MySQL.Create(&courseInfo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "添加用户基本信息失败", 
			"error": err.Error(),
		})
        return
    }

    c.JSON(http.StatusOK, gin.H{
		"status": "success", 
		"message": "用户基本信息添加成功", 
		"course": courseInfo,
	})
}

// UpdateUserBasic 更新用户基本信息信息
func UpdateUserBasicInfo(c *gin.Context) {
    var courseInfo models.UserBasicInformation
    if err := c.ShouldBindJSON(&courseInfo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
			"status": "error", 
			"message": "请求数据格式错误", 
			"error": err.Error(),
		})
        return
    }

    if err := utils.DB_MySQL.Model(&models.UserBasicInformation{}).Where("account = ?", courseInfo.Account).Updates(courseInfo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "更新用户基本信息失败", 
			"error": err.Error(),
		})
        return
    }

    c.JSON(http.StatusOK, gin.H{
		"status": "success", 
		"message": "用户基本信息信息更新成功", 
		"course": courseInfo,
	})
}

// DeleteUserBasic 删除用户基本信息信息
func DeleteUserBasicInfo(c *gin.Context) {
	courseCode := c.Param("account")

	var course models.UserBasicInformation
	if err := utils.DB_MySQL.Where("account = ?", courseCode).First(&course).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error", 
			"message": "未找到该用户基本信息", 
			"error": err.Error(),
		})
		return
	}

	if err := utils.DB_MySQL.Delete(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "删除用户基本信息失败", 
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success", 
		"message": "用户基本信息删除成功", 
		"course": course,
	})
}
