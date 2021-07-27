package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/gorilla/mux"
)

type CalcReq struct {
	Step    string `json:"Step,omitempty"`
	Product string `json:"Product,omitempty"`
	Amount  string `json:"Amount,omitempty"`
	At      string `json:"At,omitempty"`
	To      string `json:"To,omitempty"`
}

type CarbonData struct {
	Carbon float64 `json:"carbon"`
}

type EvaluateReq struct {
	Farm    string `json:"Farm,omitempty"`
	Process string `json:"Process,omitempty"`
	Retail  string `json:"Retail,omitempty"`
	Product string `json:"Product,omitempty"`
	Amount  string `json:"Amount,omitempty"`
}

type EvaluateData struct {
	Score float64 `json:"score,omitempty"`
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Data API Works!")
	fmt.Println("Endpoint Hit: homePage")
}

func Calc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: calc")

	w.Header().Set("content-type", "application/json")

	var req CalcReq
	json.NewDecoder(r.Body).Decode(&req)

	py := "python3"
	arg0 := "../scripts/single_step.py"
	arg1 := req.Step
	arg2 := req.Product
	arg3 := req.Amount
	arg4 := req.At
	arg5 := req.To
	cmd := exec.Command(py, arg0, arg1, arg2, arg3, arg4, arg5)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(stdout))

	var carbon CarbonData
	f, err := ioutil.ReadFile("../single_step.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	json.Unmarshal([]byte(f), &carbon)
	json.NewEncoder(w).Encode(carbon)
}

func Evaluate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: evaluate")

	w.Header().Set("content-type", "application/json")

	var req EvaluateReq
	json.NewDecoder(r.Body).Decode(&req)

	py := "python3"
	arg0 := "../scripts/evaluate.py"
	arg1 := req.Farm
	arg2 := req.Process
	arg3 := req.Retail
	arg4 := req.Product
	arg5 := req.Amount
	cmd := exec.Command(py, arg0, arg1, arg2, arg3, arg4, arg5)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(stdout))

	time.Sleep(50 * time.Millisecond)
	var score EvaluateData
	f, err := ioutil.ReadFile("../carbon_score.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	json.Unmarshal([]byte(f), &score)
	json.NewEncoder(w).Encode(score)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", HomePage)
	router.HandleFunc("/calc", Calc).Methods("POST")
	router.HandleFunc("/evaluate", Evaluate).Methods("POST")

	fmt.Println("Data-Proccessing API Running")
	log.Fatal(http.ListenAndServe(":8081", router))
}
