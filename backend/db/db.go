package db

import (
	"errors"
	"fmt"
	"orbit_nft/db/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ErrNoRowsAffected = errors.New("no rows affected by this operation")

type Database struct {
	DB *gorm.DB
}

func NewSQLDatabase(user, pass, host, port, dbname string) (*Database, error) {
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, dbname)
	db, err := gorm.Open(mysql.Open(URL))
	if err != nil {
		return nil, err
	}
	return &Database{
		DB: db,
	}, nil
}

func (d *Database) AutoMigrate() error {
	return d.DB.AutoMigrate(
		&model.Metadata{},
		&model.NFT{},
		&model.Marketplace{},
		&model.MakeRand{},
		&model.Hero{},
		&model.Refcode{},
	)
}

func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	sqlDB.Close()
	return nil
}
