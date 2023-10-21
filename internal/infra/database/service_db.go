package database

import (
	"github.com/NayronFerreira/api-multithreading-goexpert/config"
	"github.com/NayronFerreira/api-multithreading-goexpert/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDBConnections(appConfig config.Config) (BrasilcepDB, ViacepDB) {
	dbGorm, err := gorm.Open(sqlite.Open(appConfig.DBFilePath))
	if err != nil {
		panic(err)
	}
	dbGorm.AutoMigrate(model.ViaCepRes{}, model.BrasilCepRes{})

	dbBrasil := NewDBBrasilCep(*dbGorm)
	dbVia := NewDBViacep(*dbGorm)

	return dbBrasil, dbVia
}
