package models

type CustomerRequest struct {
	Id    int
	Name  string
	Email string
}

type ShowCustomerResponse struct {
	Id    int
	Name  string
	Email string
}

func CreateCostumer(name string, email string) CustomerRequest {

	Customer := CustomerRequest{
		Name:  name,
		Email: email,
	}

	return Customer

}

func (Customer *CustomerRequest) UpdateCostumer(id int, name, email string) {
	if name != "" {
		Customer.Name = name
	}
	if email != "" {
		Customer.Email = email
	}
}
