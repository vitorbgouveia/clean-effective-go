package usecases

import (
	"log"
	"time"

	"github.com/vitorbgouveia/clean-effective-go/internal/domain/entities"
	"github.com/vitorbgouveia/clean-effective-go/internal/domain/repositories"
	"github.com/vitorbgouveia/clean-effective-go/internal/presentation/dtos"
)

// ProcessSale process and save new sale
func ProcessSale(sale dtos.EntranceSale) dtos.OutputSale {
	log.Printf("Received a message: %v", sale)

	cashback := CalculateCashback(sale.Value)

	effectedSale := entities.Sales{
		Desc:     sale.Desc,
		Value:    sale.Value,
		Cpf:      sale.Cpf,
		Cashback: cashback,
		CreateAt: time.Now(),
	}
	repositories.SaleRepository().Save(effectedSale)

	processedSale := dtos.OutputSale{
		Cpf:      sale.Cpf,
		Cashback: cashback,
	}

	return processedSale
}
