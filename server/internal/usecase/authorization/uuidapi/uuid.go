package uuidapi

import (
	"github.com/google/uuid"
)

type UUIDProvider interface {
	CreateStringUUID() (id string)
}

type UUIDProviderImpl struct{}

func NewUUIDProvider() UUIDProvider {
	return &UUIDProviderImpl{}
}

func (hp *UUIDProviderImpl) CreateStringUUID() (id string) {
	id = uuid.New().String()

	return
}
