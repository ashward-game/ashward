package repository

import (
	"orbit_nft/db"
)

type RefcodeRepository struct {
	db *db.Database
}

func NewRefcodeRepository(db *db.Database) *RefcodeRepository {
	return &RefcodeRepository{
		db: db,
	}
}
