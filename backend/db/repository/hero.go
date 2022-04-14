package repository

import (
	"orbit_nft/db"
	"orbit_nft/db/model"
)

type HeroRepository struct {
	db *db.Database
}

func NewHeroRepository(db *db.Database) *HeroRepository {
	return &HeroRepository{
		db: db,
	}
}

func (r *HeroRepository) CreateWithOwnerGetId(owner, metadata string) (uint, error) {
	hero := &model.Hero{
		Metadata: metadata,
		Owner:    owner,
	}
	if err := r.db.DB.Model(&model.Hero{}).Create(hero).Error; err != nil {
		return 0, err
	}
	return hero.ID, nil
}

func (r *HeroRepository) FindByOwner(owner string) ([]model.Hero, error) {
	var result []model.Hero
	if err := r.db.DB.Model(&model.Hero{}).Where(&model.Hero{
		Owner: owner,
	}).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
