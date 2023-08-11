package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PgSql struct {
}

func NewPgSQL() *gorm.DB {
	log.Println("[PgSql] Connecting to DB")
	dsn := fmt.Sprint(
		"host", "=", viper.GetString("MASTER_POSTGRES_HOST"),
		" ",
		"user", "=", viper.GetString("MASTER_POSTGRES_USER"),
		" ",
		"password", "=", viper.GetString("MASTER_POSTGRES_PASSWORD"),
		" ",
		"dbname", "=", viper.GetString("MASTER_POSTGRES_DB"),
		" ",
		"port", "=", viper.GetString("MASTER_POSTGRES_PORT"),
	)
	myLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: myLogger,
	})

	if err != nil {
		log.Fatalln("[PgSql] Connect to DB Failed")
	} else {
		log.Println("[PgSql] Connect to DB Success")
	}

	return db
}
