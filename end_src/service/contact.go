package service

import (
	"jyu-service/models"
	"jyu-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetContacts 获取联系方式列表（支持分页）
func GetContacts(c *gin.Context) {
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

	var Contacts []models.ContactInformation
	var total int64

	if err := utils.DB_MySQL.Model(&models.ContactInformation{}).Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "无法获取联系方式总数", 
			"error": err.Error(),
		})
		return
	}

	if err := utils.DB_MySQL.Limit(pageSize).Offset(offset).Find(&Contacts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "无法获取联系方式数据", 
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"Contacts": Contacts,
		"total":   total,
	})
}

// AddCourse 添加联系方式信息
func AddContacts(c *gin.Context) {
    var courseInfo models.ContactInformation

    if err := c.ShouldBindJSON(&courseInfo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
			"status": "error", 
			"message": "请求数据格式错误", 
			"error": err.Error(),
		})
        return
    }

    var existingCourse models.ContactInformation
    if err := utils.DB_MySQL.Where("account = ?", courseInfo.Account).First(&existingCourse).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{
			"status": "error", 
			"message": "联系方式代码已存在",
		})
        return
    }

    if err := utils.DB_MySQL.Create(&courseInfo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "添加联系方式失败", 
			"error": err.Error(),
		})
        return
    }

    c.JSON(http.StatusOK, gin.H{
		"status": "success", 
		"message": "联系方式添加成功", 
		"course": courseInfo,
	})
}

// UpdateContacts 更新联系方式信息
func UpdateContacts(c *gin.Context) {
    var courseInfo models.ContactInformation
    if err := c.ShouldBindJSON(&courseInfo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
			"status": "error", 
			"message": "请求数据格式错误", 
			"error": err.Error(),
		})
        return
    }

    if err := utils.DB_MySQL.Model(&models.ContactInformation{}).Where("account = ?", courseInfo.Account).Updates(courseInfo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "更新联系方式失败", 
			"error": err.Error(),
		})
        return
    }

    c.JSON(http.StatusOK, gin.H{
		"status": "success", 
		"message": "联系方式信息更新成功", 
		"course": courseInfo,
	})
}

// DeleteContacts 删除联系方式信息
func DeleteContacts(c *gin.Context) {
	courseCode := c.Param("account")

	var course models.ContactInformation
	if err := utils.DB_MySQL.Where("account = ?", courseCode).First(&course).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error", 
			"message": "未找到该联系方式", 
			"error": err.Error(),
		})
		return
	}

	if err := utils.DB_MySQL.Delete(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "删除联系方式失败", 
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success", 
		"message": "联系方式删除成功", 
		"course": course,
	})
}
