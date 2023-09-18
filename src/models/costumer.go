package models

type CustomerRequest struct {
	Name  string
	Email string
}

type ShowCustomerResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
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
