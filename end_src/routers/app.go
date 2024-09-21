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


	// r.GET("/admin/GetCourses",service.GetCourses)
	// r.POST("/admin/UpdateCourses", service.UpdateCourses)
	// r.POST("/admin/UpdateCourses/:courseCode", service.UpdateCourses)
	// r.DELETE("/admin/DeleteCourses/:course_code",service.DeleteCourses)

	// admin UserAccount
	r.GET("/admin/GetUserAccount", service.GetUserAccount)
    r.PUT("/admin/UpdateUserAccount/:account", service.UpdateUserAccount)
    r.POST("/admin/AddUserAccount", service.AddUserAccount)
    r.DELETE("/admin/DeleteUserAccount/:account", service.DeleteUserAccount)

	// admin Courses
	r.GET("/admin/GetCourses", service.GetCourses)
    r.PUT("/admin/UpdateCourses/:course_code", service.UpdateCourses)
    r.POST("/admin/AddCourse", service.AddCourse)
    r.DELETE("/admin/DeleteCourses/:course_code", service.DeleteCourses)
	// admin Contacts
	r.GET("/admin/GetContacts", service.GetContacts)
    r.PUT("/admin/UpdateContacts/:account", service.UpdateContacts)
    r.POST("/admin/AddContacts", service.AddContacts)
    r.DELETE("/admin/DeleteContacts/:account", service.DeleteContacts)
	// admin UserBasicInfo
	r.GET("/admin/GetUserBasicInfo", service.GetUserBasicInfo)
    r.PUT("/admin/UpdateUserBasicInfo/:account", service.UpdateUserBasicInfo)
    r.POST("/admin/AddUserBasicInfo", service.AddUserBasicInfo)
    r.DELETE("/admin/DeleteUserBasicInfo/:account", service.DeleteUserBasicInfo)
	// admin StudentStatusInfo
	r.GET("/admin/GetStudentStatusInfo", service.GetStudentStatusInfo)
    r.PUT("/admin/UpdateStudentStatusInfo/:account", service.UpdateStudentStatusInfo)
    r.POST("/admin/AddStudentStatusInfo", service.AddStudentStatusInfo)
    r.DELETE("/admin/DeleteStudentStatusInfo/:account", service.DeleteStudentStatusInfo)

	r.GET("/admin/GetUserCourseInfomation", service.GetUserCourseInformationHandler)
	r.GET("/admin/GetUserInformation", service.GetUserInformationHandler)
	r.GET("/admin/GetContactInformation", service.GetContactInformationHandler)
	r.GET("/admin/GetStudentStatusInformation", service.GetStudentStatusInformationHandler)
	r.GET("/admin/GetGradeInformationHandler", service.GetGradeInformationHandler)

	// execl
	r.GET("/generate_excel", service.GenerateExcel)
	r.GET("/generate_excel_template", service.GenerateExcelTemplate)
	r.POST("/admin/uploadExcel", service.ImportAndGenerateAccounts)
	

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
