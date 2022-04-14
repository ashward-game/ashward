package util

import (
	"orbit_nft/testutil"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type MockUser struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func TestPaginate(t *testing.T) {
	db, teardown := testutil.NewMockDB()
	defer teardown()
	db.DB.AutoMigrate(&MockUser{})

	// insert to db
	for i := 0; i < 100; i++ {
		db.DB.Create(&MockUser{Name: "name" + strconv.Itoa(i), Email: "email" + strconv.Itoa(i)})
	}

	// select & paginate
	// this is extracted from http request
	pg := &Pagination{
		Offset: 1,
		Limit:  5,
	}

	// this is defined temporarily
	type record struct {
		MockUser `gorm:"embedded"`
		Total    int
	}

	var results []record
	db.DB.Model(&MockUser{}).Scopes(Paginate(pg)).Find(&results)

	assert.Equal(t, pg.Limit, len(results))
	assert.Equal(t, 100, results[0].Total)

	user1 := results[0].MockUser
	user5 := results[len(results)-1].MockUser
	assert.Equal(t, "name1", user1.Name)
	assert.Equal(t, "name5", user5.Name)

	// custom query: select name only
	db.DB.Model(&MockUser{}).Scopes(Paginate(pg, "name")).Find(&results)
	assert.Equal(t, pg.Limit, len(results))
	assert.Equal(t, 100, results[0].Total)

	user1 = results[0].MockUser
	user5 = results[len(results)-1].MockUser
	assert.Equal(t, "name1", user1.Name)
	assert.Equal(t, "name5", user5.Name)
	// empty email because we did not select it
	assert.Equal(t, "", user1.Email)
}
