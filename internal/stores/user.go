package stores

import (
	"github.com/cryptowatch_challenge/internal/models"
	"gorm.io/gorm"
)

func (m *UserStore) create(object *models.User) error {
	return m.Model(models.User{}).Create(object).Error
}

func (m *UserStore) getByID(userID uint32) (*models.User, bool, error) {
	var object = &models.User{}
	err := m.Model(models.User{}).Where("user_id = ?", userID).First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}

func (m *UserStore) updateMap(mm map[string]interface{}) error {
	return m.Model(models.User{}).Updates(mm).Error
}

func (m *UserStore) getByEmail(email string) (*models.User, bool, error) {
	var object = &models.User{}
	err := m.Model(models.User{}).Where("email = ?", email).First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}
