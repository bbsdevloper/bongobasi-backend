package main

import (
	"fmt"
	"net/http"

	"github.com/Prosecutor1x/citizen-connect-frontend/router"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("API for citizen-connect app")
	fmt.Println("Server stating in port 4000")
	r := router.Router()
	handler := cors.Default().Handler(r)

	http.ListenAndServe(":4000", handler)
	fmt.Println("Server started in port 4000")
}
