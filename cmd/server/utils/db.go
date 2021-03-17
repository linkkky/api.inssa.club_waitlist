package utils

import (
	"inssa_club_waitlist_backend/cmd/server/models"
	"inssa_club_waitlist_backend/configs"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var once sync.Once
var dbManager *DBManager

// DBManager is a struct to store *gorm.DB
type DBManager struct {
	Instance *gorm.DB
}

// GetDB is a singleton constructor of the DBManager
func GetDB() *DBManager {
	once.Do(func() {
		dbManager = &DBManager{}
	})
	return dbManager
}

// SetupDB setups the gorm DB configuration using the values from the environment
func (m *DBManager) SetupDB() {
	dsn := configs.Envs["PSQL_URI"]
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	for _, model := range models.GetModels() {
		db.AutoMigrate(model)
	}
	m.Instance = db
}

// InitDB is a shortcut for the GetDB().SetupDB()
func InitDB() {
	GetDB().SetupDB()
}
