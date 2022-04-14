package service

import (
	"orbit_nft/db/model"
	"orbit_nft/db/repository"

	"gorm.io/gorm"
)

type MetadataService struct {
	repo *repository.MetadataRepository
}

func NewMetadataService(r *repository.MetadataRepository) *MetadataService {
	return &MetadataService{
		repo: r,
	}
}

func (m *MetadataService) AddOrUpdateCurrentBlock(newRecord *model.Metadata) error {
	var err error
	if newRecord.ID == 0 { // no ID is set
		var foundRecord *model.Metadata
		foundRecord, err = m.repo.Find(newRecord.Name)
		if err == nil {
			newRecord.ID = foundRecord.ID
		}
	}
	switch err {
	case nil:
		err = m.repo.Update(newRecord)
		return err
	case gorm.ErrRecordNotFound:
		err = m.repo.Create(newRecord)
		return err
	default:
		return err
	}
}

func (m *MetadataService) FindMetadata(names []string) ([]model.Metadata, error) {
	metas, err := m.repo.Finds(names)
	if err != nil {
		return nil, err
	}

	return metas, nil
}
