package token

import "github.com/thimovez/service/internal/usecase"

type TokenUseCase struct {
	token usecase.TokenService
}

func New(t usecase.TokenService) *TokenUseCase {
	return &TokenUseCase{
		token: t,
	}
}

func (u *TokenUseCase) GenerateAccessToken() {

}

func (u *TokenUseCase) VerifyAccessToken() error {
	return nil
}
