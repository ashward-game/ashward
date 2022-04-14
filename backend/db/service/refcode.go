package service

import (
	"orbit_nft/db/repository"
)

type RefcodeService struct {
	repo *repository.RefcodeRepository
}

func NewRefcodeService(repo *repository.RefcodeRepository) *RefcodeService {
	return &RefcodeService{
		repo: repo,
	}
}
