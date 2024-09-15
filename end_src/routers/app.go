package routers

import (
	"jyu-service/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	//cors 配置			--- 跨域  不知道为什么要开vpn才能正确传递数据
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/admin/GetAll", service.GetAllCoursesHandler)
	r.GET("/admin/GetUserCourseInfomation", service.GetUserCourseInformationHandler)
	r.GET("/admin/GetUserInformation", service.GetUserInformationHandler)
	r.GET("/admin/GetContactInformation", service.GetContactInformationHandler)
	r.GET("/admin/GetStudentStatusInformation", service.GetStudentStatusInformationHandler)
	r.GET("/admin/GetGradeInformationHandler", service.GetGradeInformationHandler)
	r.GET("/generate_excel", service.GenerateExcel)

	r.POST("/admin/courses", service.DynamicQueryHandler)

	r.POST("/api/getChoosedNumbers", service.GetCourseChoosedNumberHandler)
	r.POST("/admin/GetAllByAccount", service.GetCourseByAccountHandle)
	r.POST("/admin/Login", service.Login)
	//insert
	r.POST("/admin/Register", service.Register)

	r.POST("/admin/EnrollCourse", service.EnrollCourseHandler)
	r.POST("/admin/UnenrollCourse", service.UnenrollCourseHandler)
	return r
}
