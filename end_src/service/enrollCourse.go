package service

import (
	"jyu-service/models"
	"jyu-service/utils"
	"net/http"
    "strconv"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// 请求体结构
type ChoosedNumbersRequest struct {
	CourseCodes []string `json:"courseCodes"`
}

// 响应体结构
type ChoosedNumbersResponse struct {
	ChoosedNumbers map[string]int `json:"choosedNumbers"`
}

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
	if err == redis.Nil {
		// 如果 Redis 中没有该课程的信息，查询 MySQL 并将数据放入 Redis
		var course models.CourseInformation
		if err := utils.DB_MySQL.Where("course_code = ?", req.CourseCode).First(&course).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "课程不存在"})
			return
		}
		maxStudentNumber = course.MaxStudentNumber

		utils.DB_Redis.Set(utils.Redis_Context, maxStudentNumberKey, maxStudentNumber, 0)
	} else if err != nil {
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

// 退课处理函数
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
func GetAllCoursesHandler(c *gin.Context) {
	var course models.CourseInformation

	courses, db := course.GetAll()

	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": db.Error.Error(),
		})
	}

	c.JSON(http.StatusOK, courses)
}

// 获取某一用户的选课记录
func GetCourseByAccountHandle(c *gin.Context) {
	var req models.UserCourse

	// 解析请求数据
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 生成 Redis 键
	redisKey := "enrollments:" + req.Account

	// 从 Redis 获取用户的选课记录
	courseCodes, err := utils.DB_Redis.SMembers(utils.Redis_Context, redisKey).Result()
	if err != nil {
		// 处理错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "从 Redis 获取选课记录失败"})
		return
	}

	if len(courseCodes) > 0 {
		// 如果 Redis 中有数据，构造 UserCourse 列表
		courses := make([]models.UserCourse, len(courseCodes))
		for i, code := range courseCodes {
			courses[i] = models.UserCourse{Account: req.Account, CourseCode: code}
		}
		c.JSON(http.StatusOK, courses)
		return
	}

	// 如果 Redis 中没有数据，从 MySQL 查询
	var userCourses []models.UserCourse
	result := utils.DB_MySQL.Where("account = ?", req.Account).Find(&userCourses)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// 将 MySQL 查询结果缓存到 Redis
	redisData := make([]string, len(userCourses))
	for i, course := range userCourses {
		redisData[i] = course.CourseCode
	}
	if err := utils.DB_Redis.SAdd(utils.Redis_Context, redisKey, redisData).Err(); err != nil {
		// 处理错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "缓存数据到 Redis 失败"})
		return
	}

	// 返回查询结果
	c.JSON(http.StatusOK, userCourses)
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

// 获取课程已选人数
func GetCourseChoosedNumberHandler(c *gin.Context) {
	var req ChoosedNumbersRequest

	// 解析请求体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	choosedNumbers := make(map[string]int)

	// 遍历 CourseCode 列表，获取每个课程的 ChoosedNumber
	for _, courseCode := range req.CourseCodes {
		var course models.CourseInformation
		course.CourseCode = courseCode

		// 调用 GetCourseChoosedNumber 函数获取已选人数
		choosedNumber, err := course.GetCourseChoosedNumber()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取已选人数失败"})
			return
		}

		choosedNumbers[courseCode] = choosedNumber
	}

	// 返回响应
	c.JSON(http.StatusOK, ChoosedNumbersResponse{
		ChoosedNumbers: choosedNumbers,
	})
}

type GradeInfoResponse struct {
	Account      string  `json:"account"`       // 学号
	CourseCode   string  `json:"course_code"`   // 课程代码
	CourseGrade  string  `json:"course_grade"`  // 成绩
	AcademicYear string  `json:"academic_year"` // 年级
	Semester     int     `json:"semester"`      // 学期
	CourseName   string  `json:"course_name"`   // 课程名称
	CourseNature int     `json:"course_nature"` // 课程性质
	Credits      float32 `json:"credits"`       // 学分
	GradePoints  float32 `json:"grade_points"`  // 绩点
}

// 查询用户成绩信息
func GetGradeInformationHandler(c *gin.Context) {
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

    // 定义一个 GradeInfoResponse 列表来存储返回的成绩信息
    var gradeInfoList []GradeInfoResponse

    // 遍历用户选的课程，根据 CourseCode 查询课程详细信息
    for _, userCourse := range arr {
        // 如果 CourseGrade 为空，则跳过这门课程
        if userCourse.CourseGrade == "" {
            continue
        }

        var course models.CourseInformation
        course.CourseCode = userCourse.CourseCode
        result := course.FindByCourseCode()
        if result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": result.Error.Error(),
            })
            return
        }

        // 计算绩点（通过学分和成绩计算）
        gradePoints := calculateGradePoints(userCourse.CourseGrade, course.Credits)

        // 构建 GradeInfoResponse 对象
        gradeInfo := GradeInfoResponse{
            Account:      req.Account,
            CourseCode:   course.CourseCode,
            CourseGrade:  userCourse.CourseGrade,
            AcademicYear: course.AcademicYear,  // 假设该字段存在于 CourseInformation 中
            Semester:     course.Semester,      // 假设该字段存在于 CourseInformation 中
            CourseName:   course.CourseName,    // 假设该字段存在于 CourseInformation 中
            CourseNature: course.CourseNature,  // 假设该字段存在于 CourseInformation 中
            Credits:      course.Credits,
            GradePoints:  gradePoints,          // 使用计算得到的绩点
        }

        // 将构建好的 GradeInfoResponse 添加到列表中
        gradeInfoList = append(gradeInfoList, gradeInfo)
    }

    // 返回构建好的成绩信息
    c.JSON(http.StatusOK, gradeInfoList)
}

// 根据100分制成绩和学分计算绩点
func calculateGradePoints(grade string, credits float32) float32 {
    // 将成绩从字符串转换为整数
    gradeScore, err := strconv.Atoi(grade)
    if err != nil {
        return 0.0 // 如果成绩格式错误，返回0.0
    }

    var gradePoint float32

    // 根据成绩分配绩点
    switch {
    case gradeScore >= 90:
        gradePoint = 4.0
    case gradeScore >= 80:
        gradePoint = 3.0
    case gradeScore >= 70:
        gradePoint = 2.0
    case gradeScore >= 60:
        gradePoint = 1.0
    default:
        gradePoint = 0.0
    }

    // 计算绩点：绩点 * 课程学分
    return gradePoint * credits
}