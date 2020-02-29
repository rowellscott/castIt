package main

import (
	"log"
	"net/http"

	"github.com/rowellscott/castIt/api"
)

func main() {
	r := api.CreateApi()
	log.Fatal(http.ListenAndServe(":8000", r))
}
