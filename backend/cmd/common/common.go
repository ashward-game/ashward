package common

import (
	"orbit_nft/db"
	"os"

	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	var logger *zap.Logger
	switch os.Getenv("ORBIT_ENV") {
	case "production":
		logger, _ = zap.NewProduction()
	default:
		logger, _ = zap.NewDevelopment()
	}
	return logger
}

func MySQL() *db.Database {
	sqlDB, err := db.NewSQLDatabase(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("MYSQL_DATABASE"))
	if err != nil {
		panic(err)
	}
	if err = sqlDB.AutoMigrate(); err != nil {
		panic(err)
	}
	return sqlDB
}
