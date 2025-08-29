package auth

import "go/courses/pkg/db"

type AuthRepository struct {
	Database *db.Db
}

func NewAuthRepository(database *db.Db) *AuthRepository {
	return &AuthRepository{
		Database: database,
	}
}

func (repo *AuthRepository) StoreAuth(auth *Auth) (*Auth, error) {
	result := repo.Database.Create(auth)

	if result.Error != nil {
		return nil, result.Error
	}

	return auth, nil
}

func (repo *AuthRepository) GetBySessionId(sessionId string) (*Auth, error) {
	var auth Auth
	result := repo.Database.First(&auth, "session_id=?", sessionId)

	if result.Error != nil {
		return nil, result.Error
	}

	return &auth, nil
}

func (repo *AuthRepository) GetByPhone(phone string) (*Auth, error) {
	var auth Auth
	result := repo.Database.First(&auth, "phone=?", phone)

	if result.Error != nil {
		return nil, result.Error
	}

	return &auth, nil
}

func (repo *AuthRepository) UpdateCodeByPhone(code int, phone string) (*Auth, error) {
	auth, err := repo.GetByPhone(phone)

	if err != nil {
		return nil, err
	}

	auth.Code = code

	result := repo.Database.Updates(auth)

	if result.Error != nil {
		return nil, result.Error
	}

	return auth, nil
}
