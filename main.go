package main

import (
	"fmt"
	"go-book-web/src/router"
	"go-book-web/src/utils"
	"log"
	"net/http"
)

func main() {
	utils.LoadTemplates()
	r := router.Generate()

	fmt.Println("Running webapp")
	log.Fatal(http.ListenAndServe(":3000", r))
}
