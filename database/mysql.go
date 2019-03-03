package database

import (
	"database/sql"
	"time"
)
import _ "github.com/go-sql-driver/mysql"
import "log"

//import "time"

var Mysql *sql.DB // MySQL

func init() {
	var err error
	Mysql, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = Mysql.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	// Connection pool and timeouts
	// 连接池 和 超时
	Mysql.SetMaxOpenConns(150) // 最大打开连接数
	Mysql.SetMaxIdleConns(10)  // 最大空闲连接数
	// 连接过期时间 测试结果如下:
	// 1如不设过期时间 连接会一直不释放 连接池内连接数量为小于等于maxopen的数字
	// 2如设置了连接过期时间
	// 2.1 连接池内连接数量在连接过期后归零
	// 2.2 如之前连接数达到了最大打开连接数 连接池内连接数会依次经历: 由maxopen => maxidle => 0
	Mysql.SetConnMaxLifetime(time.Second * 5)
}
