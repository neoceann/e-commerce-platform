// auth-service/internal/pkg/password/password.go
package password

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordService interface {
	Hash(password string) (string, error)
	Verify(hashedPassword, plainPassword string) bool
}

type bcryptService struct {
	cost int
}

func NewPasswordService(cost int) PasswordService {
	return &bcryptService{cost: cost}
}

func (s *bcryptService) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), s.cost)
	return string(bytes), err
}

func (s *bcryptService) Verify(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
