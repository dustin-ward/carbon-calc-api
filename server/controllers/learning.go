package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CalcReq struct {
	Step    string `json:"Step,omitempty"`
	Product string `json:"Product,omitempty"`
	Amount  string `json:"Amount,omitempty"`
	At      string `json:"At,omitempty"`
	To      string `json:"To,omitempty"`
}

type EvaluateReq struct {
	Farm    string `json:"Farm,omitempty"`
	Process string `json:"Process,omitempty"`
	Retail  string `json:"Retail,omitempty"`
	Product string `json:"Product,omitempty"`
	Amount  string `json:"Amount,omitempty"`
}

// Use the carbon emission calculator
//
func CarbonCalculator(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: calc")
	w.Header().Set("content-type", "application/json")

	var c CalcReq
	json.NewDecoder(r.Body).Decode(&c)
	j, _ := json.Marshal(c)
	resp, err := http.Post("http://data-processing:8081/calc", "application/json", bytes.NewBuffer(j))
	if err != nil {
		log.Fatal(err.Error())
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	w.Write(body)
}

// Gather a list of supply chains ranked by efficiency
//
func EvaluateChain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: evaluate")
	w.Header().Set("content-type", "application/json")

	var e EvaluateReq
	json.NewDecoder(r.Body).Decode(&e)
	j, _ := json.Marshal(e)
	resp, err := http.Post("http://data-processing:8081/evaluate", "application/json", bytes.NewBuffer(j))
	if err != nil {
		log.Fatal(err.Error())
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	w.Write(body)
}
