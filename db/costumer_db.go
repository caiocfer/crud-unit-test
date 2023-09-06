package db

import (
	"database/sql"
	"fmt"

	"caio.com/test-user/src/models"
)

type CostumersRepo struct {
	DB *sql.DB // Export the field by using an uppercase name
}

func CreateCostumerRepo(db *sql.DB) *CostumersRepo {
	return &CostumersRepo{db}
}

func (repo *CostumersRepo) SaveCostumerToDb(customer models.Customer) error {
	query := `INSERT INTO users (name, email) VALUES (?, ?)`
	_, err := repo.DB.Exec(query, customer.Name, customer.Email)

	if err != nil {
		fmt.Println(err)

		return err
	}

	return nil

}

func (repo *CostumersRepo) GetAllCostumers() ([]models.Customer, error) {
	query := `SELECT * from users`
	result, err := repo.DB.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	defer result.Close()

	customers := []models.Customer{}

	for result.Next() {
		var customer models.Customer

		if err = result.Scan(
			&customer.Id,
			&customer.Name,
			&customer.Email,
		); err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, err

}
