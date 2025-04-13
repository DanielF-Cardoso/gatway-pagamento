package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Name      string
	Email     string
	ApiKey    string
	Balance   float64
	mu        sync.RWMutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

func generateApiKey() string {
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	return hex.EncodeToString(randomBytes)
}

func NewAccount(name, email string) *Account {
	account := &Account{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Balance:   0,
		ApiKey:    generateApiKey(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account
}

func (a *Account) AddBalance(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
	a.UpdatedAt = time.Now()
}
