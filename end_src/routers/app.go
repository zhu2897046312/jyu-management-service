package routers

import (
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "time"
	"jyu-service/service"
)

func Router() *gin.Engine{
	r := gin.Default()
	//cors 配置			--- 跨域  不知道为什么要开vpn才能正确传递数据
	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

    r.GET("/admin/GetAll",service.GetAllCoursesHandler)
    r.GET("/admin/GetUserCourseInfomation",service.GetUserCourseInformationHandler) 
    r.GET("/admin/GetUserInformation",service.GetUserInformationHandler)
    r.POST("/admin/courses",service.DynamicQueryHandler)

    r.POST("/api/getChoosedNumbers", service.GetCourseChoosedNumberHandler)
    r.POST("/admin/GetAllByAccount",service.GetCourseByAccountHandle)
	r.POST("/admin/Login",service.Login)
    //insert
	r.POST("/admin/Register",service.Register)

    r.POST("/admin/EnrollCourse",service.EnrollCourseHandler)
    r.POST("/admin/UnenrollCourse",service.UnenrollCourseHandler)
	return r
}