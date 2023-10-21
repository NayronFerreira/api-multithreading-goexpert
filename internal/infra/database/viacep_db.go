package database

import (
	"github.com/NayronFerreira/api-multithreading-goexpert/internal/model"
	"gorm.io/gorm"
)

type ViacepDB struct {
	DB gorm.DB
}

func NewDBViacep(db gorm.DB) ViacepDB {
	return ViacepDB{DB: db}
}

func (v ViacepDB) Create(viacep model.ViaCepRes) error {
	return v.DB.Create(viacep).Error
}
