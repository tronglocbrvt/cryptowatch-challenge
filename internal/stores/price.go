package stores

import (
	"github.com/cryptowatch_challenge/internal/models"
	"gorm.io/gorm"
)

func (m *PriceStore) create(object *models.Price) error {
	return m.Model(models.Price{}).Create(object).Error
}

func (m *PriceStore) getByID(priceID uint32) (*models.Price, bool, error) {
	var object = &models.Price{}
	err := m.Model(models.Price{}).Where("price_id = ?", priceID).First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}

func (m *PriceStore) updateMap(mm map[string]interface{}) error {
	return m.Model(models.Price{}).Updates(mm).Error
}

func (m *PriceStore) getLastestPrice(marketID uint32) (*models.Price, bool, error) {
	var object = &models.Price{}
	query := m.Model(models.Price{})
	if marketID > 0 {
		query = query.Where("market_id = ?", marketID)
	}

	err := query.Last(object).Error

	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}

func (m *PriceStore) getPrices(marketID, limit, offset uint32) ([]*models.Price, error) {
	prices := make([]*models.Price, 0)

	query := m.Model(models.Price{})
	if marketID > 0 {
		query = query.Where("market_id = ?", marketID)
	}

	err := query.Limit(int(limit)).Offset(int(offset)).Error

	return prices, err
}