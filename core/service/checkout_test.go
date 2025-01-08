package service_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Mock da função ou dependência usada pela Checkout (se necessário)
type MockCheckout struct{}

func (m *MockCheckout) CreateTransaction(description string, transactionDate time.Time, transactionValue float64) (int64, error) {
	// Mock da função CreateTransaction
	return 12345, nil // Retorna um ID fictício de transação e nenhum erro
}

func TestCreateTransaction(t *testing.T) {
	description := "Test transaction"
	transactionDate := time.Now()
	transactionValue := 100.0

	checkout := &MockCheckout{}

	id, err := checkout.CreateTransaction(description, transactionDate, transactionValue)

	assert.NoError(t, err)

	assert.Greater(t, id, int64(0), "O ID da transação deveria ser maior que 0")

	assert.NotNil(t, id, "O ID da transação não pode ser nil")
}
