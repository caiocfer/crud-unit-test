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

func GetAllCostumers() []models.ShowCustomerResponse {
	conn, err := db.ConnectToDatabase()
	if err != nil {
		fmt.Println()
	}

	defer conn.Close()

	dbconn := db.CreateCostumerRepo(conn)

	costumers, err := dbconn.GetAllCostumers()
	if err != nil {
		fmt.Println(err)
	}

	return costumers
}
