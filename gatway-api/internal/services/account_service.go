package services

import (
	"github.com/DanielF-Cardoso/gatway-pagamento/gatway-api/internal/domain"
)

type AccountService struct {
	repository domain.AccountRepository
}
