package stores

import (
	"github.com/cryptowatch_challenge/internal/models"
	"gorm.io/gorm"
)

func (m *PositionStore) create(object *models.Position) error {
	return m.Model(models.Position{}).Create(object).Error
}

func (m *PositionStore) getByID(positionID uint32) (*models.Position, bool, error) {
	var object = &models.Position{}
	err := m.Model(models.Position{}).Where("position_id = ?", positionID).First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}

func (m *PositionStore) updateMap(mm map[string]interface{}) error {
	return m.Model(models.Position{}).Updates(mm).Error
}
