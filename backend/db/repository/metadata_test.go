package repository

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"orbit_nft/db/model"
	"orbit_nft/testutil"
	"testing"
)

var metaData1 = &model.Metadata{
	Name:  model.MetadataCurrentBlock,
	Value: "{value1}",
}
var metaData2 = &model.Metadata{
	Name:  model.MetadataCurrentBlock,
	Value: "{value2}",
}

func newMockMetadataRepo() (*MetadataRepository, func()) {
	db, teardown := testutil.NewMockDB()

	db.AutoMigrate()
	return NewMetadataRepository(db), teardown
}

func TestCreate(t *testing.T) {
	db, teardown := newMockMetadataRepo()
	defer teardown()
	err := db.Create(metaData1)
	assert.NoError(t, err)
}

func TestFind(t *testing.T) {
	db, teardown := newMockMetadataRepo()
	defer teardown()
	_, err := db.Find(metaData1.Name)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	err = db.Create(metaData1)
	assert.NoError(t, err)
	meta, err := db.Find(metaData1.Name)
	assert.NoError(t, err)
	assert.Equal(t, meta.Value, metaData1.Value)
}

func TestFinds(t *testing.T) {
	db, teardown := newMockMetadataRepo()
	defer teardown()
	metas, err := db.Finds([]string{metaData1.Name})
	assert.NoError(t, err)
	assert.Len(t, metas, 0)
	err = db.Create(metaData1)
	assert.NoError(t, err)
	metas, err = db.Finds([]string{metaData1.Name})
	assert.NoError(t, err)
	assert.Len(t, metas, 1)
	assert.Equal(t, metas[0].Value, metaData1.Value)
}

func TestUpdate(t *testing.T) {
	db, teardown := newMockMetadataRepo()
	defer teardown()
	err := db.Create(metaData1)
	assert.NoError(t, err)
	err = db.Update(metaData2)
	assert.NoError(t, err)
	meta, err := db.Find(metaData2.Name)
	assert.NoError(t, err)
	assert.Equal(t, meta.Value, metaData2.Value)
}
