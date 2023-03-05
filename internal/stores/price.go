package stores

import (
	"fmt"

	"github.com/cryptowatch_challenge/internal/models"
	"github.com/cryptowatch_challenge/internal/utils"
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

	err := query.Limit(int(limit)).Offset(int(offset)).Find(&prices).Error

	return prices, err
}

func (m *PriceStore) getPricesForChart(numsHour uint32, limit int) ([]*models.Price, error) {
	pricesAsc := make([]*models.Price, 0)
	pricesDesc := make([]*models.Price, 0)

	query := m.Model(models.Price{}).Where(fmt.Sprintf("created_at < (NOW() + INTERVAL '1' HOUR) AND created_at > (NOW() - INTERVAL '%d' HOUR)", numsHour))

	err := query.Limit(limit / 2).Find(&pricesAsc).Error
	if err != nil {
		return pricesAsc, nil
	}
	err = query.Limit(limit / 2).Order("price_id DESC").Find(&pricesDesc).Error
	if err != nil {
		return pricesAsc, nil
	}

	pricesDesc = utils.Reverse(pricesDesc)

	return append(pricesAsc, pricesDesc...), err
}
