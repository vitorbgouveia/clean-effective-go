package usecases

import (
	"github.com/vitorbgouveia/clean-effective-go/internal/domain/repositories"
	"github.com/vitorbgouveia/clean-effective-go/internal/presentation/dtos"
)

// SumSaleValues return resume values of cpf
func SumSaleValues(p dtos.ParamsGetSumSaleValues) dtos.SumSaleValues {
	var sales = repositories.SaleRepository().FindByCPF(p.Cpf)

	var s dtos.SumSaleValues

	for _, sale := range sales {
		s.TotalValue += sale.Value
		s.TotalCashback += sale.Cashback
	}

	return s
}
