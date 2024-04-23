package controllers

import (
	"go-book-web/src/utils"
	"net/http"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func ShowUserRegistrationPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "user-register.html", nil)
}
