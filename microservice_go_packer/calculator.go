package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "encoding/json"
	"log"
    "strconv"
)

var add = "+"
var sub = "-"
var div = "/"
var mult = "*"

type Calculator struct {
    FirstNumber float64 
    SecondNumber float64 
    Operation string 
    Result float64 
}

var history []Calculator

func Sum(responseWriter http.ResponseWriter, request *http.Request) {
    param := mux.Vars(request)
    firstNumber, firstError := strconv.ParseFloat(param["firstNumber"], 64)
    secondNumber, secondError := strconv.ParseFloat(param["secondNumber"], 64)

    if firstError != nil {
		http.Error(responseWriter, firstError.Error(), http.StatusInternalServerError)
	} else if secondError != nil {
		http.Error(responseWriter, secondError.Error(), http.StatusInternalServerError)
	} else {
        result := firstNumber + secondNumber
        calculator := Calculator{firstNumber, secondNumber, add, result} 
        AddHistory(calculator)

        response, err := json.Marshal(calculator)
        if err != nil {
            http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
        } else {
            responseWriter.Header().Set("Content-Type", "application/json")
            responseWriter.Write(response)
        }
    }
}

func Sub(responseWriter http.ResponseWriter, request *http.Request)  {
    param := mux.Vars(request)
    firstNumber, firstError := strconv.ParseFloat(param["firstNumber"], 64)
    secondNumber, secondError := strconv.ParseFloat(param["secondNumber"], 64)

    if firstError != nil {
        http.Error(responseWriter, firstError.Error(), http.StatusInternalServerError)
    } else if secondError != nil {
        http.Error(responseWriter, secondError.Error(), http.StatusInternalServerError)
    } else {
        result := firstNumber - secondNumber
        calculator := Calculator{firstNumber, secondNumber, sub, result}
        AddHistory(calculator)

        response, err := json.Marshal(calculator)
        if err != nil {
            http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
        } else {
            responseWriter.Header().Set("Content-Type", "application/json")
            responseWriter.Write(response)
        }
    }
}

func Mult(responseWriter http.ResponseWriter, request *http.Request) {
    param := mux.Vars(request)
    firstNumber, firstError := strconv.ParseFloat(param["firstNumber"], 64)
    secondNumber, secondError := strconv.ParseFloat(param["secondNumber"], 64)

    if firstError != nil {
        http.Error(responseWriter, firstError.Error(), http.StatusInternalServerError)
    } else if secondError != nil {
        http.Error(responseWriter, secondError.Error(), http.StatusInternalServerError)
    } else {
        result := firstNumber * secondNumber
        calculator := Calculator{firstNumber, secondNumber, mult, result}
        AddHistory(calculator)

        response, err := json.Marshal(calculator)
        if err != nil {
            http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
        } else {
            responseWriter.Header().Set("Content-Type", "application/json")
            responseWriter.Write(response)
        }
    }
}

func Div(responseWriter http.ResponseWriter, request *http.Request) {
    param := mux.Vars(request)
    firstNumber, firstError := strconv.ParseFloat(param["firstNumber"], 64)
    secondNumber, secondError := strconv.ParseFloat(param["secondNumber"], 64)

    if firstError != nil {
        http.Error(responseWriter, firstError.Error(), http.StatusInternalServerError)
    } else if secondError != nil {
        http.Error(responseWriter, secondError.Error(), http.StatusInternalServerError)
    } else {
        result := firstNumber / secondNumber
        calculator := Calculator{firstNumber, secondNumber, div, result}
        AddHistory(calculator)

        response, err := json.Marshal(calculator)
        if err != nil {
            http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
        } else {
            responseWriter.Header().Set("Content-Type", "application/json")
            responseWriter.Write(response)
        }
    }
}

func History(responseWriter http.ResponseWriter, request *http.Request) {
    response, err := json.Marshal(history)
    if err != nil {
        http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
        return
    }
    responseWriter.Header().Set("Content-Type", "application/json")
    responseWriter.Write(response)
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
