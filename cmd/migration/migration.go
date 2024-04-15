package migration

import (
	"context"
	"fmt"

	"kindercastle_backend/internal/app/config"
	migrationSQL "kindercastle_backend/internal/pkg/dbase/migration/mysql"
)

func Migrate(ctx context.Context, sqldsn string) {
	if sqldsn == "" {
		conf := config.GetConfig()
		mysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName)
		migrationSQL.MySqlMigration(mysqlDSN)
	} else {
		migrationSQL.MySqlMigration(sqldsn)
	}
}
