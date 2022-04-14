package service

import (
	"bytes"
	"orbit_nft/crypto"
	"orbit_nft/db/model"
	"orbit_nft/db/repository"
	"orbit_nft/testutil"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
)

func TestHashAlreadyExists(t *testing.T) {
	db, teardown := testutil.NewMockDB()
	defer teardown()
	repo := repository.NewMakeRandRepository(db)
	service := NewMakeRandService(repo)

	hash := "0x123abc"
	random := "0xabc123"

	err := service.Commit(hash, random)
	assert.NoError(t, err)

	err = service.Commit(hash, random)
	assert.Error(t, err)
}

func TestCommit(t *testing.T) {
	db, teardown := testutil.NewMockDB()
	defer teardown()
	repo := repository.NewMakeRandRepository(db)
	service := NewMakeRandService(repo)

	hash := "0x123abc"
	random := "0xabc123"
	err := service.Commit(hash, random)
	assert.NoError(t, err)

	var actual model.MakeRand
	err = db.DB.Model(&model.MakeRand{
		Hash: hash,
	}).First(&actual).Error
	assert.NoError(t, err)
	assert.Equal(t, actual.Hash, hash)
	assert.Equal(t, actual.Random, random)
}

func TestHashNotExistsAndCallReveal(t *testing.T) {
	db, teardown := testutil.NewMockDB()
	defer teardown()
	repo := repository.NewMakeRandRepository(db)
	service := NewMakeRandService(repo)

	hash := "0x123abc"

	_, err := service.Reveal(hash)
	assert.Error(t, err)
}

func TestReveal(t *testing.T) {
	db, teardown := testutil.NewMockDB()
	defer teardown()
	repo := repository.NewMakeRandRepository(db)
	service := NewMakeRandService(repo)

	hash := "0x123abc"

	rand, err := crypto.MakeRand()
	assert.NoError(t, err)

	random := hexutil.Encode(rand)

	err = service.Commit(hash, random)
	assert.NoError(t, err)

	err = db.DB.Model(&model.MakeRand{
		Hash: hash,
	}).First(&model.MakeRand{}).Error
	assert.NoError(t, err)

	actual, err := service.Reveal(hash)
	assert.NoError(t, err)
	b := bytes.Equal(actual, rand)
	assert.True(t, b)
}
