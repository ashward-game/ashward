package repository

import (
	"orbit_nft/db/model"
	"orbit_nft/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newMockMakeRandRepo() (*MakeRandRepository, func()) {
	db, teardown := testutil.NewMockDB()

	db.AutoMigrate()
	return NewMakeRandRepository(db), teardown
}

func TestCreateMakeRand(t *testing.T) {
	db, teardown := newMockMakeRandRepo()
	defer teardown()

	hash := "0x123abc"

	random := "0xbac123"

	err := db.Create(hash, random)
	assert.NoError(t, err)

	var actual model.MakeRand
	err = db.db.DB.Model(&model.MakeRand{
		Hash: hash,
	}).First(&actual).Error
	assert.NoError(t, err)

	assert.Equal(t, actual.Hash, hash)
	assert.Equal(t, actual.Random, random)
}

func TestGetByHash(t *testing.T) {
	db, teardown := newMockMakeRandRepo()
	defer teardown()

	hash := "0x123abc"

	random := "0xbac123"

	err := db.Create(hash, random)
	assert.NoError(t, err)

	actual, err := db.Get(hash)
	assert.NoError(t, err)
	assert.Equal(t, actual.Hash, hash)
	assert.Equal(t, actual.Random, random)
}
