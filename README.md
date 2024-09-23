# 信息管理平台

```
学习内容
1. 使用 docker run 直接设置允许远程访问并配置密码：
    docker run --name my-redis -p 6379:6379 -d redis redis-server --requirepass yourpassword --bind 0.0.0.0 --protected-mode no
参数解释：
--requirepass yourpassword：指定 Redis 的访问密码。
--bind 0.0.0.0：允许外部（远程）主机访问 Redis。如果不设置此项，Redis 只会监听 localhost，无法进行远程访问。
--protected-mode no：关闭保护模式。Redis 默认在保护模式下运行以阻止远程访问，将其关闭允许外部连接。
```

```
                UserAccount - account （not modifined Unique）                                      
                    |   
        admins     students     teachers  
                    
admins: PUT GET UPDATE ->   studentsInfo\teachersInfo     

studentsInfo includes: accountInfo userBasicInfo contectsInfo statusInfo enrollCourseInfo(Lists)  gradesInfoByCourse(Lists)
teachersInfo includes: accountInfo userBasicInfo contectsInfo teachCourseInfo(Lists) gradesInfoByCourse(Lists)

                    coursesInfo - courseCode  （Unique）                               gradesInfo

```