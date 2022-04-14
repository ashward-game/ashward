package testutil

import (
	"log"
	"orbit_nft/db"
	"orbit_nft/db/model"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMockDB() (*db.Database, func()) {
	sqlDB, err := gorm.Open(sqlite.Open("./orbit_test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("[gorm open] %s", err)
	}

	teardown := func() {
		db, _ := sqlDB.DB()
		db.Close()
		os.Remove("./orbit_test.db")
	}

	err = sqlDB.AutoMigrate(
		&model.Metadata{},
		&model.NFT{},
		&model.Marketplace{},
		&model.MakeRand{},
		&model.Hero{},
		&model.Refcode{},
	)
	if err != nil {
		teardown()
		log.Fatal(err)
	}

	return &db.Database{DB: sqlDB}, teardown
}
