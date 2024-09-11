package service

import (
	"jyu-service/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// 选课处理函数
func EnrollCourseHandler(c *gin.Context) {
	var req models.UserCourse
	// 解析请求数据
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用插入选课数据的函数
	if err := req.EnrollCourse(req.Account, req.CourseCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "选课失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "选课成功"})
}

func UnenrollCourseHandler(c *gin.Context) {
	var req models.UserCourse

	// 解析请求数据
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用 UnenrollCourse 函数执行退课操作
	if result := req.UnenrollCourse(); result.Error != nil {
		// 如果退课操作出错
		c.JSON(http.StatusInternalServerError, gin.H{"error": "退课失败"})
		return
	} else if result.RowsAffected == 0 {
		// 如果没有找到符合条件的记录
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到对应的选课记录"})
		return
	}

	// 返回退课成功信息
	c.JSON(http.StatusOK, gin.H{"message": "退课成功"})
}

// 选课列表
func GetAllCoursesHandler(c *gin.Context)	{
	var course models.CourseInformation

	courses,db := course.GetAll()

	if db.Error!= nil{
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": db.Error.Error(),
        })
    }

	c.JSON(http.StatusOK, courses)
}

// 获取某一用户的选课记录
func GetCourseByAccountHandle(c *gin.Context){
	var req models.UserCourse

	// 解析请求数据
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	arr , db := req.GetByAccount(req.Account)
	if db.Error!= nil{
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": db.Error.Error(),
        })
    }

	c.JSON(http.StatusOK, arr)
}

func DynamicQueryHandler(c *gin.Context) {
	var conditions map[string]interface{}
	var courses models.CourseInformation
	if err := c.BindJSON(&conditions); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}
	results, err := courses.DynamicQuery(conditions)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to execute query"})
		return
	}
	c.JSON(200, results)
}

// 查询用户选课信息
func GetUserCourseInformationHandler(c *gin.Context) {
	// 获取查询参数中的 account 值
    account := c.Query("account")
    
    if account == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "缺少 account 参数"})
        return
    }

    // 根据学号查询选课信息
    var req models.UserCourse
    req.Account = account

    arr, db := req.GetByAccount(req.Account)
    if db.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": db.Error.Error(),
        })
        return
    }

	// 遍历用户选的课程，根据 CourseCode 查询课程详细信息
    var courses []models.CourseInformation
	for _, userCourse := range arr {
        var course models.CourseInformation
		course.CourseCode = userCourse.CourseCode
        result := course.FindByCourseCode()
        if result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": result.Error.Error(),
            })
            return
        }
        courses = append(courses, course)
    }

    // 返回查询到的选课信息
    c.JSON(http.StatusOK, courses)
}