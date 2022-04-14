package service

import (
	"orbit_nft/db/repository"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type MakeRandService struct {
	repo *repository.MakeRandRepository
}

func NewMakeRandService(repo *repository.MakeRandRepository) *MakeRandService {
	return &MakeRandService{
		repo: repo,
	}
}

func (s *MakeRandService) Commit(hash, random string) error {
	return s.repo.Create(hash, random)
}

func (s *MakeRandService) Reveal(hash string) ([]byte, error) {
	result, err := s.repo.Get(hash)
	if err != nil {
		return nil, err
	}

	random, err := hexutil.Decode(result.Random)
	if err != nil {
		return nil, err
	}

	return random, nil
}
