package stores

import (
	"github.com/cryptowatch_challenge/internal/models"
	"gorm.io/gorm"
)

type PriceStore struct {
	*gorm.DB
}

func NewPriceStore(db *gorm.DB) *PriceStore {
	return &PriceStore{db}
}

func (m *PriceStore) Create(object *models.Price) error {
	object.BeforeCreate()
	return m.create(object)
}

func (m *PriceStore) GetByID(id uint32) (*models.Price, bool, error) {
	return m.getByID(id)
}

func (m *PriceStore) UpdateMap(mm map[string]interface{}) error {
	return m.updateMap(mm)
}

func (m *PriceStore) GetLastestPrice(marketID uint32) (*models.Price, bool, error) {
	return m.getLastestPrice(marketID)
}

func (m *PriceStore) GetPrices(marketID, limit, offset uint32) ([]*models.Price, error) {
	return m.getPrices(marketID, limit, offset)
}

func (m *PriceStore) GetPricesForChart(numsHour uint32, limit int) ([]*models.Price, error) {
	return m.getPricesForChart(numsHour, limit)
}
