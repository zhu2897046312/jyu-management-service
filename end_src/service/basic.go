package service

import (
	"jyu-service/models"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){
	var user models.UserAccount
	//接收json数据
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(300, gin.H{
            "error": err.Error(),
        })
        return
	}
	fmt.Printf("user.UserName: %v\n", user.Account)
	fmt.Printf("user.Password: %v\n", user.Password)
	if user.Account == "" || user.Password == ""{
		c.JSON(300, gin.H{
            "error": "用户名或密码不能为空",
        })
        return
	}

	//查询数据库
	usr_after_find := models.UserAccount{
		Account: user.Account,
	}
	db := user.FindByAccount(&usr_after_find)
	if db.Error != nil{
		c.JSON(300, gin.H{
            "error": db.Error.Error(),
        })
        return
	}

	// 检查是否找到了数据
    if db.RowsAffected <= 0 {
		c.JSON(300, gin.H{
            "msg": "没有找到数据",
        })
        return 
    } 
	
	// TODO: md5加密

	//密码比对
	if usr_after_find.Password != user.Password{
        c.JSON(300, gin.H{
            "error": "密码错误",
        })
        return
    }
    c.JSON(200, gin.H{
        "message": "登录成功",
		"data": usr_after_find.ChatType,	//是否是管理员
    })

	// TODO: 发token
}

func Register(c *gin.Context){
	user := models.UserAccount{}
	//接收json数据
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(300, gin.H{
            "error": err.Error(),
        })
        return
	}
	fmt.Printf("user.UserName: %v\n", user.Account)
	fmt.Printf("user.Password: %v\n", user.Password)
	fmt.Printf("user.IsAdministrator: %v\n", user.ChatType)
	//查询数据库
	db := user.FindByAccount(&user)

	//没找到
	if db.RowsAffected == 0{
		// TODO: md5加密

		//插入数据库
		db = user.Insert(&user)
		if db.Error!= nil{
			c.JSON(300, gin.H{
				"error": db.Error.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "注册成功",
		})
		return
	}else{
		//找到了
		c.JSON(300, gin.H{
			"message": "用户已存在",
		})
		return
	}
}

func GetUserInformationHandler(c *gin.Context){
	// 获取查询参数中的 account 值
    account := c.Query("account")
    
    if account == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "缺少 account 参数"})
        return
    }

	var userInformation models.UserBasicInformation
	userInformation.Account = account
	// 查询数据库
	db := userInformation.FindByAccount(&userInformation)
    if db.Error!= nil{
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": db.Error.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, userInformation)
}