package database

import (
	"github.com/NayronFerreira/api-multithreading-goexpert/internal/model"
	"gorm.io/gorm"
)

type BrasilcepDB struct {
	DB gorm.DB
}

func NewDBBrasilCep(db gorm.DB) BrasilcepDB {
	return BrasilcepDB{DB: db}
}

func (b BrasilcepDB) Create(brasilcep model.BrasilCepRes) error {
	return b.DB.Create(brasilcep).Error
}
