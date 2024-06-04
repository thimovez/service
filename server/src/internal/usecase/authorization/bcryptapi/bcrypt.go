package bcryptapi

import "golang.org/x/crypto/bcrypt"

type BcryptProvider interface {
	HashPassword(password string) (hashedPassword []byte, err error)
	ComparePassword(hashedPassword, password []byte) error
}

type BcryptProviderImpl struct{}

func NewBcryptProvider() BcryptProvider {
	return &BcryptProviderImpl{}
}

func (bp *BcryptProviderImpl) HashPassword(password string) (hashedPassword []byte, err error) {
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	return
}

func (bp *BcryptProviderImpl) ComparePassword(hashedPassword, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return err
	}

	return nil
}
