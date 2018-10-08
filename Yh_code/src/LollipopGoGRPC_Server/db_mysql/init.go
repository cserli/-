package dbif

import (
	"database/sql"
	"fmt"
	"glog-master"
	"runtime"
	//	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 初始化操作
var db *sql.DB

// 数据操作  线程池的使用操作
// 短信验证接口 微信认证接口操作
func init() {
	db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

// 获取链接指针函数
func getMySQL() *sql.DB {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			strerr := fmt.Sprintf("%s", err)
			glog.Info("清除奖券定时器：", strerr)
		}
	}()
	if MySQLPool == nil {
		MySQLPool = make(chan *sql.DB, MAX_POOL_SIZE)
	}
	//  --  数据操作
	if len(MySQLPool) == 0 {
		go func() {
			for i := 0; i < MAX_POOL_SIZE/2; i++ {
				mysql := new(sql.DB)
				var err error
				var StrConnection = ""
				StrConnection = "root" + ":" + "ruilide2016" + "@tcp(" + "120.24.219.60" + ":3306)/" + "gl_RuiLiDe"
				mysql, err = sql.Open("mysql", StrConnection)
				if err != nil {
					glog.Info("Connect Fail!")
					continue
				}
				putMySQL(mysql)
			}

			// 清除并发连接
			runtime.Goexit()
		}()
	}
	return <-MySQLPool
}
