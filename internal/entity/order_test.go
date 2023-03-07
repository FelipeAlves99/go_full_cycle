package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// First word must be Test - original way of GO
// func TestIfItGetAnErrorIfIdIsBlank(t *testing.T) {
// 	order := Order{}
// 	err := order.Validate()

// 	if err == nil {
// 		t.Error("Expected error, got nil")
// 	}
// }

func TestIfItGetAnErrorIfIdIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "invalid id")
}

func TestIfItGetAnErrorIfPriceIsBlank(t *testing.T) {
	order := Order{Id: "1"}
	assert.Error(t, order.Validate(), "invalid price")
}

func TestIfItGetAnErrorIfTaxIsBlank(t *testing.T) {
	order := Order{Id: "2", Price: 10.0}
	assert.Error(t, order.Validate(), "invalid tax")
}

func TestWithAllValidParams(t *testing.T) {
	order := Order{Id: "3", Price: 10.0, Tax: 1}
	assert.NoError(t, order.Validate(), "invalid id")
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}
