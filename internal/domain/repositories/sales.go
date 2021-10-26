package repositories

import (
	"context"

	"github.com/vitorbgouveia/clean-effective-go/configs/database"
	"github.com/vitorbgouveia/clean-effective-go/internal/domain/entities"
	"github.com/vitorbgouveia/clean-effective-go/tools"
	"go.mongodb.org/mongo-driver/bson"
)

type Saler interface {
	FindByCPF(cpf string) []entities.Sales
	Save(sale entities.Sales)
}

type sale struct{}

func SaleRepository() Saler {
	var s Saler = sale{}
	return s
}

var clientDB, ctxDB = database.DB().Client()

// FindByCPF return all sales by cpf
func (sl sale) FindByCPF(cpf string) []entities.Sales {
	collection := clientDB.Database("cashback").Collection("sales")

	ctx, cancel := context.WithCancel(ctxDB)
	defer cancel()

	sales, err := collection.Find(ctx, bson.M{"cpf": cpf})
	tools.FatalError(err, "Error to load sales")

	s := make([]entities.Sales, 0)
	errAll := sales.All(ctx, &s)
	tools.FatalError(errAll, "Error to decode sales")

	return s
}

// SaveSale insert new sale on database
func (sl sale) Save(sale entities.Sales) {
	collection := clientDB.Database("cashback").Collection("sales")

	ctx, cancel := context.WithCancel(ctxDB)
	defer cancel()

	_, err := collection.InsertOne(ctx, sale)
	tools.FatalError(err, "Error to save effected sale")
}

