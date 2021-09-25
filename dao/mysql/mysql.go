package mysql

import (
	"fmt"
	"web_app/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var DB *sqlx.DB

func InitDB(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
	if DB, err = sqlx.Connect("mysql", dsn); err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	DB.SetMaxOpenConns(cfg.MaxOpenConns)
	DB.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

func Close() {
	_ = DB.Close()
}
