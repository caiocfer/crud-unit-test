package models

import (
	"testing"
)

func TestCreateCostumer(t *testing.T) {
	customer := CustomerRequest{
		Name:  "John Doe",
		Email: "john@doe.com",
	}

	testCostumer := CreateCostumer(
		customer.Name,
		customer.Email,
	)

	expectedCostumer := CustomerRequest{
		Name:  customer.Name,
		Email: customer.Email,
	}

	if testCostumer != expectedCostumer {
		t.Errorf("Test case failed: Expecter %v but got %v", expectedCostumer, testCostumer)
	}

}
