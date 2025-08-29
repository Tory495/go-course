package auth

type AuthService struct {
	AuthRepository *AuthRepository
}

func NewAuthService(repo *AuthRepository) *AuthService {
	return &AuthService{
		AuthRepository: repo,
	}
}
