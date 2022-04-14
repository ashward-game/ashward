package service

import (
	"errors"
	"math/big"
	"orbit_nft/db/model"
	"orbit_nft/db/repository"
	"orbit_nft/db/util"

	"gorm.io/gorm"
)

type MarketplaceService struct {
	marketplaceRepo *repository.MarketplaceRepository
	nftRepo         *repository.NftRepository
}

func NewMarketplaceService(repo *repository.MarketplaceRepository, nftRepo *repository.NftRepository) *MarketplaceService {
	return &MarketplaceService{
		marketplaceRepo: repo,
		nftRepo:         nftRepo,
	}
}

func (s *MarketplaceService) OpenOffer(tokenId uint, seller string, price *big.Int) error {
	nft, err := s.nftRepo.FindByTokenID(tokenId).FirstToken()
	if err != nil {
		return err
	}

	if nft.Owner != seller {
		return errors.New(ErrOwnerAndSellerNotMatch)
	}

	if nft.IsForSale() {
		return errors.New(ErrNftIsOnMarketplace)
	}

	marketplace := model.NewMarketplace(nft.ID, nft.Owner, model.OnSale, price)

	return s.marketplaceRepo.Transaction(func(tx *gorm.DB) error {
		if err := s.marketplaceRepo.WithTx(tx).Create(marketplace); err != nil {
			return err
		}

		nft.Sell(price)
		if err := s.nftRepo.WithTx(tx).Update(nft); err != nil {
			return err
		}
		return nil
	})
}

func (s *MarketplaceService) CancelOffer(tokenId uint, seller string) error {
	nft, err := s.nftRepo.FindByTokenID(tokenId).FirstToken()
	if err != nil {
		return err
	}

	if nft.Owner != seller {
		return errors.New(ErrOwnerAndSellerNotMatch)
	}

	if !nft.IsForSale() {
		return errors.New(ErrNftNotOnMarketplace)
	}

	marketplace, err := s.marketplaceRepo.FindOnSaleByTokenId(nft.ID)
	if err != nil {
		return err
	}

	return s.marketplaceRepo.Transaction(func(tx *gorm.DB) error {
		marketplace.Status = model.Cancelled
		if err := s.marketplaceRepo.WithTx(tx).Update(marketplace); err != nil {
			return err
		}

		nft.NotForSale()
		if err := s.nftRepo.WithTx(tx).Update(nft); err != nil {
			return err
		}
		return nil
	})
}

func (s *MarketplaceService) Purchase(tokenId uint, buyer string) error {
	nft, err := s.nftRepo.FindByTokenID(tokenId).FirstToken()
	if err != nil {
		return err
	}

	if !nft.IsForSale() {
		return errors.New(ErrNftNotOnMarketplace)
	}

	marketplace, err := s.marketplaceRepo.FindOnSaleByTokenId(nft.ID)
	if err != nil {
		return err
	}

	return s.marketplaceRepo.Transaction(func(tx *gorm.DB) error {
		marketplace.Status = model.Sold
		marketplace.Buyer = buyer
		err = s.marketplaceRepo.Update(marketplace)
		if err != nil {
			return err
		}

		nft.NotForSale()
		nft.Owner = buyer
		err = s.nftRepo.Update(nft)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *MarketplaceService) GetTradingOfAddressWithPg(pg *util.Pagination, address string) (int64, []map[string]interface{}, error) {
	// sanity check
	var err error
	var total int64

	result, err := s.marketplaceRepo.FindTradingHistoryOf(pg, address)
	if err != nil {
		return 0, make([]map[string]interface{}, 0), err
	}

	if len(result) == 0 {
		return 0, make([]map[string]interface{}, 0), nil
	}

	var actualResults []map[string]interface{}

	total = result[0]["total"].(int64)
	for _, v := range result {
		delete(v, "total")
		v["token_id"] = v["token_id"].(int64)
		actualResults = append(actualResults, v)
	}

	return total, actualResults, nil
}
