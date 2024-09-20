package service

import (
	"jyu-service/models"
	"jyu-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUserBasic 获取用户基本信息列表（支持分页）
func GetUserAccount(c *gin.Context) {
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

	var UserAccount []models.UserAccount
	var total int64

	if err := utils.DB_MySQL.Model(&models.UserAccount{}).Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "无法获取用户基本信息总数", 
			"error": err.Error(),
		})
		return
	}

	if err := utils.DB_MySQL.Limit(pageSize).Offset(offset).Find(&UserAccount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error", 
			"message": "无法获取用户基本信息数据", 
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"UserAccount": UserAccount,
		"total":   total,
	})
}

// AddCourse 添加用户基本信息信息
func AddUserAccount(c *gin.Context) {
	var userAccount models.UserAccount

    // 1. 从请求中解析用户信息（不包含账号）
    if err := c.ShouldBindJSON(&userAccount); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request payload",
        })
        return
    }

    // 3. 插入数据库
    newAccount,db := userAccount.Insert_auto()
	if db.Error!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to create user account",
        })
        return
	}

    // 4. 返回成功响应
    c.JSON(http.StatusOK, gin.H{
        "message": "User account created successfully",
        "account": newAccount,
    })
}

// UpdateUserBasic 更新用户基本信息信息
func UpdateUserAccount(c *gin.Context) {
    var courseInfo models.UserAccount
    if err := c.ShouldBindJSON(&courseInfo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
			"status": "error", 
			"message": "请求数据格式错误", 
			"error": err.Error(),
		})
        return
    }

    if err := utils.DB_MySQL.Model(&models.UserAccount{}).Where("account = ?", courseInfo.Account).Updates(courseInfo).Error; err != nil {
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
func DeleteUserAccount(c *gin.Context) {
	courseCode := c.Param("account")

	var course models.UserAccount
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
