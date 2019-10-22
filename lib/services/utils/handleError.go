package utils

import (
	"encoding/json"
	"net/http"

	"github.com/YaroslavRozum/go-boilerplate/lib/errors"
)

func HandleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	if _, ok := err.(*errors.Error); !ok {
		w.Write([]byte(`{"status":0, "reason":"Server Error" }`))
		return
	}
	jsonError, _ := json.Marshal(err)
	w.Write(jsonError)
	return
}
