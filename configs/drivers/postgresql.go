package drivers

import (
	"fmt"
	"log"
	"os"
	"time"

	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgreSql struct {
}

func (d *PostgreSql) Connect(host string, port int, user string, password string, dbname string, debug bool) *gorm.DB {
	var db *gorm.DB
	var err error

	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbname, port)
	if debug {
		db, err = gorm.Open(driver.Open(conn), &gorm.Config{
			SkipDefaultTransaction: true,
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,
					LogLevel:      logger.Info,
					Colorful:      false,
				},
			),
		})
	} else {
		db, err = gorm.Open(driver.Open(conn), &gorm.Config{
			SkipDefaultTransaction: true,
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: 200 * time.Millisecond,
					LogLevel:      logger.Warn,
					Colorful:      false,
				},
			),
		})
	}

	if err != nil {
		log.Printf("Gorm PostgreSQL: %+v \n", err)
		panic(err)
	}

	return db
}
