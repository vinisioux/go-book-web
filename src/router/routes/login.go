package routes

import (
	"go-book-web/src/controllers"
	"net/http"
)

var loginRoutes = []Route{
	{
		URI:          "/",
		Method:       http.MethodGet,
		Function:     controllers.LoadLoginPage,
		AuthRequired: false,
	},
	{
		URI:          "/login",
		Method:       http.MethodGet,
		Function:     controllers.LoadLoginPage,
		AuthRequired: false,
	},
}
