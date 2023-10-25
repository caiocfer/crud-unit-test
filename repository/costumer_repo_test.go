package repository

import (
	"testing"

	"caio.com/test-user/src/models"
)

func TestCreateCostumer(t *testing.T) {

	var mockCreateCostumer = CreateCostumer
	testCustomer := models.CustomerRequest{

		Name:  "John Doe",
		Email: "john@doe.com",
	}
	expectedCustomer := models.CustomerRequest{
		Name:  "",
		Email: "",
	}

	tests := []struct {
		name             string
		mockFunc         func()
		expectedCustomer models.CustomerRequest
	}{
		{
			name: "Success creating customer",
			mockFunc: func() {
				mockCreateCostumer = func(name, email string) {
					expectedCustomer.Name = "John Doe"
					expectedCustomer.Email = "john@doe.com"

				}
			},
		},
	}

	originalCreateCostumer := mockCreateCostumer

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockFunc()
			mockCreateCostumer(expectedCustomer.Name, expectedCustomer.Email)
			if expectedCustomer.Name != testCustomer.Name {
				t.Errorf("Error Name: expected %v but got %v", testCustomer.Name, expectedCustomer.Name)
			}
			if expectedCustomer.Email != testCustomer.Email {
				t.Errorf("Error Email: expected %v but got %v", testCustomer.Email, expectedCustomer.Email)
			}
		})

	}

	mockCreateCostumer = originalCreateCostumer

}

func TestGetAllCostumers(t *testing.T) {
	var mockGetAllCostumers = GetAllCostumers
	customerList := []models.ShowCustomerResponse{
		{
			Id:    0,
			Name:  "John Doe",
			Email: "john@doe.com",
		},
		{
			Id:    1,
			Name:  "Jane Doe",
			Email: "jane@doe.com",
		},
	}

	tests := []struct {
		name                 string
		mockFunc             func()
		expectedCustomerList []models.ShowCustomerResponse
	}{
		{
			name: "Success retrieving users",
			mockFunc: func() {
				mockGetAllCostumers = func() ([]models.ShowCustomerResponse, error) {
					return []models.ShowCustomerResponse{
						{
							Id:    0,
							Name:  "John Doe",
							Email: "john@doe.com",
						},
						{
							Id:    1,
							Name:  "Jane Doe",
							Email: "jane@doe.com",
						},
					}, nil
				}
			},
			expectedCustomerList: customerList,
		},
	}
	for _, testCases := range tests {
		t.Run(testCases.name, func(t *testing.T) {
			testCases.mockFunc()
			testCustomers, _ := mockGetAllCostumers()
			for i := range testCustomers {
				if testCustomers[i] != testCases.expectedCustomerList[i] {
					t.Errorf("Expected customer %+v, but got %+v", testCustomers[i], testCases.expectedCustomerList[i])
				}
			}

		})

	}
}

func TestGetCustomerByID(t *testing.T) {
	var mockGetCustomerByID = GetCustomerByID

	tests := []struct {
		name             string
		customerID       int
		mockFunc         func()
		expectedCustomer models.ShowCustomerResponse
	}{
		{
			name:       "Success finding id",
			customerID: 0,
			mockFunc: func() {

				mockGetCustomerByID = func(id int) (models.ShowCustomerResponse, error) {
					return models.ShowCustomerResponse{
						Id:    0,
						Name:  "John Doe",
						Email: "john@Doe",
					}, nil
				}
			},
			expectedCustomer: models.ShowCustomerResponse{
				Id:    0,
				Name:  "John Doe",
				Email: "john@Doe",
			},
		},
	}

	originalGetCustomerByID := mockGetCustomerByID

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockFunc()
			testCustomer, _ := mockGetCustomerByID(testCase.customerID)

			if testCustomer != testCase.expectedCustomer {
				t.Errorf("Error: expected %v but got %v", testCustomer, testCase.expectedCustomer)
			}

		})
	}

	mockGetCustomerByID = originalGetCustomerByID

}
