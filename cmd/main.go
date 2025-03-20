package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/willeei/pfa-go/internal/order/infra/database"
	orderUsecase "github.com/willeei/pfa-go/internal/order/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	usecase := orderUsecase.NewCalculateFinalPriceUseCase(repository)

	input := orderUsecase.OrderInputDTO{
		ID:    uuid.NewString(),
		Price: 10,
		Tax:   1,
	}
	output, err := usecase.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The Final Price is: %f\n", output.FinalPrice)
}
