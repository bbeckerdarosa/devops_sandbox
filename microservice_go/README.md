## Microservice Calculator Golang

### Instructions

- To run the command below, you need install the *golang*

- Enter in the root directory microservice_go/ and type in the terminal:

```
    go run calculator.go
```

- You can open the browser and access *localhost:8080*

### Available Endpoints

#### Sum

```
    http://localhost:8080/calc/sum/{firstNumber}/{secondNumber}
```

#### Sub

```
    http://localhost:8080/calc/sub/{firstNumber}/{secondNumber}
```

#### Div

```
    http://localhost:8080/calc/div/{firstNumber}/{secondNumber}
```

#### Mult

```
    http://localhost:8080/calc/mult/{firstNumber}/{secondNumber}
```

#### History

```
    http://localhost:8080/calc/history
```