package main

import (
	"net/http"
)

func main() {
	r := createRouter()
	http.ListenAndServe(":3000", r)
}
