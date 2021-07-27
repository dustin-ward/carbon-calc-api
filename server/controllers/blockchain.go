package controllers

import (
	"fmt"
	"net/http"
)

// Blockchain stuff here
//
func Temp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: temp")
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(`{ "message": "not implemented yet" }`))
}
