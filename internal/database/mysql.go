package database

import (
	"fmt"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

const dsnFormat = "%s:%s@%s/%s?charset=utf8&parseTime=True"

func connect(tracer *tracesdk.TracerProvider) (*gorm.DB, error) {
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	address := os.Getenv("DATABASE_ADDRESS")
	database := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf(dsnFormat, user, password, address, database)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,  // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // smart configure based on used version
	}), &gorm.Config{})
	if err != nil {
		return db, err
	}

	err = db.Use(otelgorm.NewPlugin(otelgorm.WithDBName("product-db"), otelgorm.WithTracerProvider(tracer)))
	return db, err

}
