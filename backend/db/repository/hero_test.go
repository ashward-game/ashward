package repository

import (
	"orbit_nft/db/model"
	"orbit_nft/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateHero(t *testing.T) {
	db, teardown := testutil.NewMockDB()
	defer teardown()

	repo := NewHeroRepository(db)

	metadata := "metadata"
	owner := "0x123"
	id, err := repo.CreateWithOwnerGetId(owner, metadata)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), id)

	var actual model.Hero
	err = db.DB.Model(&model.Hero{}).Where(&model.Model{
		ID: id,
	}).First(&actual).Error
	assert.NoError(t, err)
	assert.Equal(t, owner, actual.Owner)
	assert.Equal(t, metadata, actual.Metadata)

	metadata2 := "metadata2"
	owner2 := "0x456"
	id2, err := repo.CreateWithOwnerGetId(owner2, metadata2)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), id)

	var actual2 model.Hero
	err = db.DB.Model(&model.Hero{}).Where(&model.Model{
		ID: id2,
	}).First(&actual2).Error
	assert.NoError(t, err)
	assert.Equal(t, owner2, actual2.Owner)
	assert.Equal(t, metadata2, actual2.Metadata)

	var actual3 []model.Hero
	err = db.DB.Model(&model.Hero{}).Find(&actual3).Error
	assert.NoError(t, err)
	assert.Equal(t, 2, len(actual3))
}

func TestFindHeroByOwner(t *testing.T) {
	db, teardown := testutil.NewMockDB()
	defer teardown()

	repo := NewHeroRepository(db)

	metadata := "metadata"
	metadata2 := "metadata2"
	owner := "0x123"
	owner2 := "0x456"
	id1, err := repo.CreateWithOwnerGetId(owner, metadata)
	assert.NoError(t, err)

	id2, err := repo.CreateWithOwnerGetId(owner2, metadata)
	assert.NoError(t, err)

	id3, err := repo.CreateWithOwnerGetId(owner, metadata2)
	assert.NoError(t, err)

	actual, err := repo.FindByOwner(owner)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(actual))
	assert.Equal(t, id1, actual[0].ID)
	assert.Equal(t, id3, actual[1].ID)

	actual2, err := repo.FindByOwner(owner2)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(actual2))
	assert.Equal(t, id2, actual2[0].ID)

	actual3, err := repo.FindByOwner("0x789")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(actual3))
}
