package test

import (
	"inssa_club_waitlist_backend/cmd/server/utils"
	"testing"

	"github.com/tj/assert"
)

func TestDBSingleton(t *testing.T) {
	db1 := utils.GetDB()
	db2 := utils.GetDB()
	utils.GetDB().SetupDB()

	assert.Equal(t, true, db1.Instance == db2.Instance)
}
