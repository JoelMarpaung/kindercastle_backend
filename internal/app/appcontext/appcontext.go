package appcontext

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"kindercastle_backend/internal/app/config"
)

type AppContext struct {
	conf *config.Config
}

func NewAppContext(cfg *config.Config) AppContext {
	return AppContext{
		conf: cfg,
	}
}

func (ac AppContext) GetDBConnection() *sqlx.DB {
	dbDSN := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", ac.conf.DBUser, ac.conf.DBPass, ac.conf.DBHost, ac.conf.DBPort, ac.conf.DBName)

	db, err := sqlx.Open("mysql", dbDSN)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Hour)
	db.SetMaxOpenConns(ac.conf.DBMaxOpen)
	db.SetMaxIdleConns(ac.conf.DBMaxIdle)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
