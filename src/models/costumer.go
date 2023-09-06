package models

type Customer struct {
	Id    int
	Name  string
	Email string
}

func CreateCostumer(id int, name string, email string) Customer {

	Customer := Customer{
		Id:    id,
		Name:  name,
		Email: email,
	}

	return Customer

}

func (Customer *Customer) UpdateCostumer(id int, name, email string) {
	if name != "" {
		Customer.Name = name
	}
	if email != "" {
		Customer.Email = email
	}
}
