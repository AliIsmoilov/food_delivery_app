package db

import (
	"fmt"
	"log"
	"monorepo/src/order_service/configs"
	"os"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Initialize database connection then connect with postgres
func Init(config *configs.Configuration) (*gorm.DB, error) {

	fmt.Println("init db")
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.PostgresUser, config.PostgresPassword, config.PostgresHost, config.PostgresPort, config.PostgresDatabase)
	// fmt.Println("11111", dbURL)
	// // fmt.Println("22222", config.PostgresPassword)
	// // fmt.Println("33333", config.PostgresHost)
	// // fmt.Println("44444", config.PostgresPort)
	// // fmt.Println("555555", config.PostgresDatabase)

	// m, err := migrate.New("file://src/restaurant_service/migrations", dbURL)
	// if err != nil {
	// 	log.Fatal("error in creating migrations: ", zap.Error(err))
	// }
	// fmt.Printf("")
	// if err = m.Up(); err != nil {
	// 	log.Println("error updating migrations: ", zap.Error(err))
	// }

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second * 1, // Slow SQL threshold
			// LogLevel:                  logger.Info,                                               // For debugging you can set Info level
			LogLevel: logger.Error, // For debugging you can set Info level
			// IgnoreRecordNotFoundError: false,        // Ignore ErrRecordNotFound error for logger
			// ParameterizedQueries:      false,        // Don't include params in the SQL log
			Colorful: true, // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Failed to get database", err)
	} else {
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetMaxIdleConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	return db, nil
}
