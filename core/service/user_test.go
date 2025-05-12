package service_test

import (
	"errors"
	"testing"

	coreErrors "github.com/luancpereira/APICheckout/core/errors"
	"github.com/luancpereira/APICheckout/core/service"
	"github.com/stretchr/testify/assert"
)

// MockCrypt implementa a interface de hashing e verificação de senha
type MockCrypt struct{}

func (MockCrypt) MakeHash(password string) (string, error) {
	if password == "error" {
		return "", errors.New("hash error")
	}
	return "hashedPassword", nil
}

func (MockCrypt) Check(password, hash string) error {
	if password != "correctPassword" {
		return errors.New("invalid password")
	}
	return nil
}

// MockUserService é uma estrutura para testes
type MockUserService struct {
	service.User
}

// Teste para validações de senha
func TestValidatePassword(t *testing.T) {
	user := service.User{}

	t.Run("Senhas diferentes", func(t *testing.T) {
		err := user.ValidatePassword("Senha123!", "Senha456!")
		coreErr, ok := err.(*coreErrors.CoreError)
		assert.True(t, ok)
		assert.Equal(t, "error.public.user.password.mismatch", coreErr.Key)
	})

	t.Run("Senha curta", func(t *testing.T) {
		err := user.ValidatePassword("A!", "A!")
		coreErr, ok := err.(*coreErrors.CoreError)
		assert.True(t, ok)
		assert.Equal(t, "error.public.user.password.size", coreErr.Key)
	})

	t.Run("Sem maiúsculas suficientes", func(t *testing.T) {
		err := user.ValidatePassword("a!a!a!aa", "a!a!a!aa")
		coreErr, ok := err.(*coreErrors.CoreError)
		assert.True(t, ok)
		assert.Equal(t, "error.public.user.password.uppers", coreErr.Key)
	})

	t.Run("Sem caracteres especiais suficientes", func(t *testing.T) {
		err := user.ValidatePassword("AAaaaaaa", "AAaaaaaa")
		coreErr, ok := err.(*coreErrors.CoreError)
		assert.True(t, ok)
		assert.Equal(t, "error.public.user.password.special.characters", coreErr.Key)
	})

	t.Run("Senha válida", func(t *testing.T) {
		err := user.ValidatePassword("AA!!aa11", "AA!!aa11")
		assert.NoError(t, err)
	})
}

func (m *MockUserService) GetIDByEmail(email string) (int64, error) {
	if email == "valido@teste.com" {
		return 1, nil
	}
	return 0, nil
}

func TestValidateEmail(t *testing.T) {
	user := &MockUserService{}

	t.Run("E-mail inválido", func(t *testing.T) {
		err := user.ValidateEmail("invalido-email", 0)
		coreErr, ok := err.(*coreErrors.CoreError)
		assert.True(t, ok)
		assert.Equal(t, "error.validation.email.invalid", coreErr.Key)
	})
}
