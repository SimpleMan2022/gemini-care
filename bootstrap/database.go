package bootstrap

import (
	"fmt"
	postgres2 "gemini-care/external/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase(env *ENV) {
	var (
		DB_HOST      = env.DB_HOST
		DB_PORT      = env.DB_PORT
		DB_USER      = env.DB_USER
		DB_PASSWORD  = env.DB_PASSWORD
		DB_NAME      = env.DB_NAME
		DB_SSL_MODE  = env.DB_SSL_MODE
		DB_TIME_ZONE = env.DB_TIME_ZONE
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT, DB_SSL_MODE, DB_TIME_ZONE)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	MigrateDatabase(db)
}

func MigrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&postgres2.User{},
		&postgres2.History{})
}
