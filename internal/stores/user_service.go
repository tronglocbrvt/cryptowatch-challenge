package stores

import (
	"github.com/cryptowatch_challenge/internal/models"
	"gorm.io/gorm"
)

type UserStore struct {
	*gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db}
}

func (m *UserStore) Create(object *models.User) error {
	object.BeforeCreate()
	return m.create(object)
}

func (m *UserStore) GetByID(id uint32) (*models.User, bool, error) {
	return m.getByID(id)
}

func (m *UserStore) UpdateMap(mm map[string]interface{}) error {
	return m.updateMap(mm)
}

func (m *UserStore) GetByEmail(email string) (*models.User, bool, error) {
	return m.getByEmail(email)
}
