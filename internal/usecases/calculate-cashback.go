package usecases

import (
	"github.com/vitorbgouveia/clean-effective-go/configs/env"
)

// CalculateCashback return value cashback
func CalculateCashback(value float32) float32 {
	cashbackPercent := env.Envs.Domain.CashbackPercent

	cashback := value * cashbackPercent / 100
	return cashback
}
