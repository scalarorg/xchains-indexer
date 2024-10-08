package db

import (
	"errors"

	"github.com/scalarorg/xchains-indexer/db/models"
	"gorm.io/gorm"
)

func FindOrCreateDenomByBase(db *gorm.DB, base string) (models.Denom, error) {
	if base == "" {
		return models.Denom{}, errors.New("base is required")
	}

	denom := models.Denom{
		Base: base,
	}
	err := db.Where(&denom).FirstOrCreate(&denom).Error
	return denom, err
}

func FindOrCreateAddressByAddress(db *gorm.DB, address string) (models.Address, error) {
	if address == "" {
		return models.Address{}, errors.New("address is required")
	}

	addr := models.Address{
		Address: address,
	}
	err := db.Where(&addr).FirstOrCreate(&addr).Error
	return addr, err
}

func GetChains(db *gorm.DB) ([]models.Chain, error) {
	var chains []models.Chain
	if err := db.Find(&chains).Error; err != nil {
		return nil, err
	}
	return chains, nil
}
