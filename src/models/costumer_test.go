package models

import (
	"testing"
)

func TestCreateCostumer(t *testing.T) {
	customer := Customer{
		Id:    0,
		Name:  "John Doe",
		Email: "john@doe.com",
	}

	testCostumer := CreateCostumer(
		customer.Id,
		customer.Name,
		customer.Email,
	)

	expectedCostumer := Customer{
		Id:    customer.Id,
		Name:  customer.Name,
		Email: customer.Email,
	}

	if testCostumer != expectedCostumer {
		t.Errorf("Test case failed: Expecter %v but got %v", expectedCostumer, testCostumer)
	}

}
