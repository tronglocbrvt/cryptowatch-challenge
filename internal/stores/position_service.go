package stores

import (
	"github.com/cryptowatch_challenge/internal/models"
	"gorm.io/gorm"
)

type PositionStore struct {
	*gorm.DB
}

func NewPositionStore(db *gorm.DB) *PositionStore {
	return &PositionStore{db}
}

func (m *PositionStore) Create(object *models.Position) error {
	object.BeforeCreate()
	return m.create(object)
}

func (m *PositionStore) GetByID(id uint32) (*models.Position, bool, error) {
	return m.getByID(id)
}

func (m *PositionStore) UpdateMap(mm map[string]interface{}) error {
	return m.updateMap(mm)
}
