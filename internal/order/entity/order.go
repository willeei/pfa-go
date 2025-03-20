package entity

import "errors"

type OrderRepositoryInterface interface {
	Save(order *Order) error
}

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price, tax float64) (*Order, error) {
	order := Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}

func (o Order) IsValid() error {
	if o.ID == "" {
		return errors.New("the order ID is required")
	}
	if o.Price == 0 {
		return errors.New("the order Price is required")
	}
	if o.Tax == 0 {
		return errors.New("the order Tax is required")
	}
	return nil
}
