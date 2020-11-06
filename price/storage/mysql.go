package storage

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gtforge/isavinof/test/config"
)

var MySQLInstance sql.DB

func init() {
	host := config.ConfigInstance.GetMysqlConfig().Addr
	name := fmt.Sprintf("%v:%v@%v",
		config.ConfigInstance.GetMysqlConfig().User,
		config.ConfigInstance.GetMysqlConfig().Passwd,
		config.ConfigInstance.GetMysqlConfig().DBName,
	)
	db, err := sql.Open(host, name)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
