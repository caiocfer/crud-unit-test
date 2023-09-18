package repository

import (
	"fmt"

	"caio.com/test-user/db"
	"caio.com/test-user/src/models"
)

func CreateCostumer(name string, email string) {

	costumer := models.CreateCostumer(name, email)

	conn, err := db.ConnectToDatabase()
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	dbconn := db.CreateCostumerRepo(conn)

	dbconn.SaveCostumerToDb(costumer)

}

func GetAllCostumers() ([]models.ShowCustomerResponse, error) {
	conn, err := db.ConnectToDatabase()
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	dbconn := db.CreateCostumerRepo(conn)

	customer, err := dbconn.GetAllCostumers()
	if err != nil {
		fmt.Println(err)
	}

	return customer, err
}

func GetCustomerByID(id int) (models.ShowCustomerResponse, error) {
	conn, err := db.ConnectToDatabase()
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	dbconn := db.CreateCostumerRepo(conn)

	customer, err := dbconn.GetCustomerByID(id)
	if err != nil {
		fmt.Println(err)
	}

	return customer, err
}
