package db

import (
	"testing"

	"caio.com/test-user/src/models"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestSaveCostumerToDb(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}

	defer db.Close()

	costumerRepo := CreateCostumerRepo(db)

	testCostumer := models.CustomerRequest{
		Name:  "John",
		Email: "john@doe.com",
	}
	mockQuery := `INSERT INTO users \(name, email\) VALUES \(\?, \?\)`
	mock.ExpectExec(mockQuery).
		WithArgs(testCostumer.Name, testCostumer.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = costumerRepo.SaveCostumerToDb(testCostumer)

	if err != nil {
		t.Errorf("TestSaveCostumerToDB customer failed, got error %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectation: %v", err)
	}

}

func TestGetAllCostumers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	costumersRepo := CreateCostumerRepo(db)

	mockQuery := `SELECT \* from users`

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).
		AddRow(1, "John Doe", "john@doe.com").
		AddRow(2, "Jane Doe", "jane@doe.com")

	mock.ExpectQuery(mockQuery).WillReturnRows(rows)

	customers, err := costumersRepo.GetAllCostumers()
	if err != nil {
		t.Errorf("Expected no erro, but got: %v", err)
	}

	expectedCustomers := []models.ShowCustomerResponse{
		{Id: 1, Name: "John Doe", Email: "john@doe.com"},
		{Id: 2, Name: "Jane Doe", Email: "jane@doe.com"},
	}

	if len(customers) != len(expectedCustomers) {
		t.Errorf("Expected %d customers, but got %d", len(expectedCustomers), len(customers))
	}

	for i := range expectedCustomers {
		if customers[i] != expectedCustomers[i] {
			t.Errorf("Expected customer %+v, but got %+v", expectedCustomers, customers)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %s", err)
		}
	}

}
