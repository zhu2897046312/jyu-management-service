package service

import (
	"jyu-service/models"
	"jyu-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// 选课处理函数
func EnrollCourseHandler(c *gin.Context) {
	var req models.UserCourse

    //1. 解析请求数据
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
        return
    }	
    
	// Redis 中缓存的课程信息 key
	courseKey := "course:" + req.CourseCode
    choosedNumberKey := courseKey + ":choosed_number"
    maxStudentNumberKey := courseKey + ":max_student_number"

	// 获取 Redis 中的已选人数
    choosedNumber, err := utils.DB_Redis.Get(utils.Redis_Context, choosedNumberKey).Int()
    if err == redis.Nil {
        // 如果 Redis 中没有该课程的信息，查询 MySQL 并将数据放入 Redis
        var course models.CourseInformation
        if err := utils.DB_MySQL.Where("course_code = ?", req.CourseCode).First(&course).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "课程不存在"})
            return
        }
        choosedNumber = course.ChoosedNumber
        maxStudentNumber := course.MaxStudentNumber

        // 将数据存入 Redis
        utils.DB_Redis.Set(utils.Redis_Context, choosedNumberKey, choosedNumber, 0)
        utils.DB_Redis.Set(utils.Redis_Context, maxStudentNumberKey, maxStudentNumber, 0)
    } else if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis 读取失败"})
        return
    }

	// 获取 Redis 中的最大人数
    maxStudentNumber, err := utils.DB_Redis.Get(utils.Redis_Context, maxStudentNumberKey).Int()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis 读取失败"})
        return
    }

	// 检查是否超过人数上限
    if choosedNumber >= maxStudentNumber {
        c.JSON(http.StatusBadRequest, gin.H{"error": "课程人数已满"})
        return
    }

	// 更新 Redis 中的已选人数
    utils.DB_Redis.Incr(utils.Redis_Context, choosedNumberKey)

	// 插入选课记录到 Redis
    enrollmentKey := "enrollments:" + req.Account
    utils.DB_Redis.SAdd(utils.Redis_Context, enrollmentKey, req.CourseCode)

	c.JSON(http.StatusOK, gin.H{"message": "选课成功"})
}

func UnenrollCourseHandler(c *gin.Context) {
	var req models.UserCourse

    // 解析请求数据
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
        return
    }

    // 从 Redis 中移除选课记录
    enrollmentKey := "enrollments:" + req.Account
    utils.DB_Redis.SRem(utils.Redis_Context, enrollmentKey, req.CourseCode)

    // 更新 Redis 中的已选人数
    courseKey := "course:" + req.CourseCode
    choosedNumberKey := courseKey + ":choosed_number"
    utils.DB_Redis.Decr(utils.Redis_Context, choosedNumberKey)

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