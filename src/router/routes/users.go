package routes

import (
	"go-book-web/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:          "/register",
		Method:       http.MethodGet,
		Function:     controllers.ShowUserRegistrationPage,
		AuthRequired: false,
	},
}
