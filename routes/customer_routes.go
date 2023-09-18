package routes

import "caio.com/test-user/src/controller"

var customerRoute = []Route{
	{
		URI:      "/customers",
		Function: controller.GetAllCostumers,
		Method:   "GET",
	},
}
