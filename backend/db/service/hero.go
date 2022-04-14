package service

import (
	"orbit_nft/db/model"
	"orbit_nft/db/repository"
)

type HeroService struct {
	repo *repository.HeroRepository
}

func NewHeroService(repo *repository.HeroRepository) *HeroService {
	return &HeroService{
		repo: repo,
	}
}

func (r *HeroService) CreateWithOwnerGetId(owner, metadata string) (uint, error) {
	return r.repo.CreateWithOwnerGetId(owner, metadata)
}

func (r *HeroService) FindByOwner(owner string) ([]model.Hero, error) {
	return r.repo.FindByOwner(owner)
}
