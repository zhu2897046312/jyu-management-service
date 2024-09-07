package test

import (
	"fmt"
	"jyu-service/utils"
	"path/filepath"
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestGetMysqlConfig(t *testing.T) {
	configPath := filepath.Join(".config")
    configFileName := "app"

	dns := utils.GetMysqlConfig(configPath, configFileName)

	// 打印 DSN 值
    fmt.Printf("dsn: %v\n", dns)

	t.Logf("dsn: %v", dns)
}

func TestInitMysql(t *testing.T) {
    utils.InitMysql()

    // 断言数据库连接是否成功
    assert.NotNil(t, utils.DB_MySQL, "数据库连接失败，DB_MySQL 为空")

    // 测试数据库简单查询
    var result int
    err := utils.DB_MySQL.Raw("SELECT 1").Scan(&result).Error
    assert.NoError(t, err, "执行简单查询失败")
    assert.Equal(t, 1, result, "简单查询结果不匹配")
}
