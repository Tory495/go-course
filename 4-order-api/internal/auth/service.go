package auth

import "go/courses/pkg/generator"

type AuthService struct {
	AuthRepository *AuthRepository
}

func NewAuthService(repo *AuthRepository) *AuthService {
	return &AuthService{
		AuthRepository: repo,
	}
}

func (service *AuthService) SendSms(phone string) error {
	code, err := generator.GenerateSmsCode()

	if err != nil {
		return err
	}

	service.AuthRepository.UpdateCodeByPhone(code, phone)
	return nil
}
