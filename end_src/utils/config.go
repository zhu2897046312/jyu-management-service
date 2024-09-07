package utils


import (
	"github.com/spf13/viper"
	"log"
	"fmt"
	"net/url"
)

// MySQLConfig 定义 MySQL 的配置结构
type MySQLConfig struct {
    User      string `mapstructure:"user"`
    Password  string `mapstructure:"password"`
    Host      string `mapstructure:"host"`
    Port      int    `mapstructure:"port"`
    Database  string `mapstructure:"database"`
    Charset   string `mapstructure:"charset"`
    ParseTime bool   `mapstructure:"parseTime"`
    Loc       string `mapstructure:"loc"`
}

// RedisConfig 定义 Redis 的配置结构
type RedisConfig struct {
    Addr        string `mapstructure:"addr"`
    Password    string `mapstructure:"password"`
    DB          int    `mapstructure:"DB"`
    PoolSize    int    `mapstructure:"poolSize"`
    MinIdleConn int    `mapstructure:"minIdleConn"`
}

// Config 包含 MySQL 和 Redis 的配置
type Config struct {
    MySQL MySQLConfig `mapstructure:"mysql"`
    Redis RedisConfig `mapstructure:"redis"`
}

// 构建 MySQL DNS 连接字符串
func (c MySQLConfig) MysqlDSN() string {
    loc := url.QueryEscape(c.Loc)
    return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
        c.User, c.Password, c.Host, c.Port, c.Database, c.Charset, c.ParseTime, loc)
}

func GetMysqlConfig(path string, fileName string) (string) {
	// 初始化 Viper
    viper.SetConfigName(fileName) // 配置文件名
    viper.AddConfigPath(path)    // 配置文件路径

    // 读取配置文件
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("读取配置文件失败: %v", err)
    }

    var config Config

    // 解析配置到结构体
    if err := viper.Unmarshal(&config); err != nil {
        log.Fatalf("配置文件解析失败: %v", err)
    }

	return config.MySQL.MysqlDSN()
}