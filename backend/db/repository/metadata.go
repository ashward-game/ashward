package repository

import (
	"orbit_nft/db"
	"orbit_nft/db/model"
)

type MetadataRepository struct {
	db *db.Database
}

func NewMetadataRepository(db *db.Database) *MetadataRepository {
	return &MetadataRepository{
		db: db,
	}
}

func (e *MetadataRepository) Create(meta *model.Metadata) error {
	return e.db.DB.Create(meta).Error
}

func (e *MetadataRepository) Update(meta *model.Metadata) error {
	if meta.ID != 0 {
		return e.db.DB.Model(meta).Updates(meta).Error
	} else {
		return e.db.DB.Where(&model.Metadata{Name: meta.Name}).Updates(meta).Error
	}
}

func (e *MetadataRepository) Find(name string) (*model.Metadata, error) {
	var meta model.Metadata
	if err := e.db.DB.Where(&model.Metadata{Name: name}).Take(&meta).Error; err != nil {
		return nil, err
	}
	return &meta, nil
}

func (e *MetadataRepository) Finds(names []string) ([]model.Metadata, error) {
	var metas []model.Metadata
	if err := e.db.DB.Where("name in ?", names).Find(&metas).Error; err != nil {
		return nil, err
	}
	return metas, nil
}
