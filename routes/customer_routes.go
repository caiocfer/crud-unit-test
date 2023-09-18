package routes

import "caio.com/test-user/src/controller"

var customerRoute = []Route{
	{
		URI:      "/customers",
		Function: controller.GetAllCostumers,
		Method:   "GET",
	},
	{
		URI:      "/customers/{id}",
		Function: controller.GetCustomerByID,
		Method:   "GET",
	},
}
