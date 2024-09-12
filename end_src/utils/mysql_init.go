package utils

import (
	"log"
	"os"
	"time"
	"path/filepath"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB_MySQL *gorm.DB
)

func init(){
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	configPath := filepath.Join("config")
    configFileName := "app"

	dns := GetMysqlConfig(configPath, configFileName)

	db, err := gorm.Open(mysql.Open(dns),
		&gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	DB_MySQL = db
}

func InitMysql(){
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	log.Println(viper.GetString("mysql.dns"))

	configPath := filepath.Join("e:\\WorkSpace\\jyu-management-service\\end_src\\config")
    configFileName := "app"

	dns := GetMysqlConfig(configPath, configFileName)

	db, err := gorm.Open(mysql.Open(dns),
		&gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	DB_MySQL = db
}



