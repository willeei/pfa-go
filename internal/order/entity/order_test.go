package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willeei/pfa-go/internal/order/entity"
)

func TestGivenAnEmptyId_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{}
	assert.EqualError(t, order.IsValid(), "the order ID is required")
}

func TestGivenAnInvalidPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{ID: "123"}
	assert.EqualError(t, order.IsValid(), "the order Price is required")
}

func TestGivenAnInvalidTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := entity.Order{ID: "123", Price: 10}
	assert.EqualError(t, order.IsValid(), "the order Tax is required")
}

func TestGivenAValidParams_WhenCallNewOrder_ThenShouldCreateANewOrderWithAllParams(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 1)
	assert.NoError(t, err)
	assert.Equal(t, order.ID, "123")
	assert.Equal(t, order.Price, 10.0)
	assert.Equal(t, order.Tax, 1.0)
}

func TestGivenAValidParams_WhenCallCalculateFinalPrice_ThenShouldCalculateFinalPriceAndSetItOnFinalPriceProperty(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 1)
	assert.NoError(t, err)
	err = order.CalculateFinalPrice()
	assert.NoError(t, err)
	assert.Equal(t, order.FinalPrice, 11.0)
}
