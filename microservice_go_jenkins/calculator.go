package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var add = "+"
var sub = "-"
var div = "/"
var mult = "*"

type Calculator struct {
	FirstNumber  float64
	SecondNumber float64
	Operation    string
	Result       float64
}

var history []Calculator

func Sum(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	firstNumber, firstError := strconv.ParseFloat(param["firstNumber"], 64)
	secondNumber, secondError := strconv.ParseFloat(param["secondNumber"], 64)

	if firstError != nil {
		http.Error(w, firstError.Error(), http.StatusInternalServerError)
	} else if secondError != nil {
		http.Error(w, secondError.Error(), http.StatusInternalServerError)
	} else {
		result := firstNumber + secondNumber
		calculator := Calculator{firstNumber, secondNumber, add, result}
		AddHistory(calculator)

		response, err := json.Marshal(calculator)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		}
	}
}

func Sub(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	firstNumber, firstError := strconv.ParseFloat(param["firstNumber"], 64)
	secondNumber, secondError := strconv.ParseFloat(param["secondNumber"], 64)

	if firstError != nil {
		http.Error(w, firstError.Error(), http.StatusInternalServerError)
	} else if secondError != nil {
		http.Error(w, secondError.Error(), http.StatusInternalServerError)
	} else {
		result := firstNumber - secondNumber
		calculator := Calculator{firstNumber, secondNumber, sub, result}
		AddHistory(calculator)

		response, err := json.Marshal(calculator)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		}
	}
}

func Mult(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	firstNumber, firstError := strconv.ParseFloat(param["firstNumber"], 64)
	secondNumber, secondError := strconv.ParseFloat(param["secondNumber"], 64)

	if firstError != nil {
		http.Error(w, firstError.Error(), http.StatusInternalServerError)
	} else if secondError != nil {
		http.Error(w, secondError.Error(), http.StatusInternalServerError)
	} else {
		result := firstNumber * secondNumber
		calculator := Calculator{firstNumber, secondNumber, mult, result}
		AddHistory(calculator)

		response, err := json.Marshal(calculator)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		}
	}
}

func Div(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	firstNumber, firstError := strconv.ParseFloat(param["firstNumber"], 64)
	secondNumber, secondError := strconv.ParseFloat(param["secondNumber"], 64)

	if firstError != nil {
		http.Error(w, firstError.Error(), http.StatusInternalServerError)
	} else if secondError != nil {
		http.Error(w, secondError.Error(), http.StatusInternalServerError)
	} else {
		result := firstNumber / secondNumber
		calculator := Calculator{firstNumber, secondNumber, div, result}
		AddHistory(calculator)

		response, err := json.Marshal(calculator)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		}
	}
}

func History(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(history)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func AddHistory(calculator Calculator) {
	history = append(history, calculator)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/calc/sum/{firstNumber}/{secondNumber}", Sum).Methods("GET")
	router.HandleFunc("/calc/sub/{firstNumber}/{secondNumber}", Sub).Methods("GET")
	router.HandleFunc("/calc/mult/{firstNumber}/{secondNumber}", Mult).Methods("GET")
	router.HandleFunc("/calc/div/{firstNumber}/{secondNumber}", Div).Methods("GET")
	router.HandleFunc("/calc/history", History).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
