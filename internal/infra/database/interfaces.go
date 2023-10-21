package database

import "github.com/NayronFerreira/api-multithreading-goexpert/internal/model"

type ViaCepInterface interface {
	Create(viacepRes model.ViaCepRes) error
}

type BrasilCepInterface interface {
	Create(brasilcepRes model.BrasilCepRes) error
}
