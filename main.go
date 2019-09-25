package main

import (
	"net/http"
)

func main() {
	r := createRouter()
	http.ListenAndServe(":3000", r)
}

type L struct {
	name string
}

func (l *L) setName(name string) {
	l.name = name
}

type Xl struct {
	L *L
}
