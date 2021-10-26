package dtos

// EntranceSale are the entry params for the sale
type EntranceSale struct {
	Desc  string  `json:"desc"`
	Value float32 `json:"value"`
	Cpf   string  `json:"cpf"`
}

// OutputSale are the sale output params
type OutputSale struct {
	Cpf      string  `json:"cpf"`
	Cashback float32 `json:"cashback"`
}

// SumSaleValues is params to return resume sum sale values
type SumSaleValues struct {
	TotalValue    float32 `json:"totalValue"`
	TotalCashback float32 `json:"totalCashback"`
}

// ParamsGetSumSaleValues is params to get resume totalize values
type ParamsGetSumSaleValues struct {
	Cpf string `json:"cpf"`
}
