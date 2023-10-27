package helpers

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type HelperProvider interface {
	CreateStringUUID() (id string)
	HashPassword(password string) (hashedPassword []byte, err error)
}

type HelperProviderImpl struct{}

func NewHelperProvider() HelperProvider {
	return &HelperProviderImpl{}
}

func (hp *HelperProviderImpl) CreateStringUUID() (id string) {
	id = uuid.New().String()

	return
}

func (hp *HelperProviderImpl) HashPassword(password string) (hashedPassword []byte, err error) {
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	return
}
