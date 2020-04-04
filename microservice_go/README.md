## Microservice Calculator Golang

### Instruções

- Para rodar o comando a seguir, você precisa instalar o *golang*

- Entre na raiz onde está o script go e digite pelo terminal:

```
    go run calculator.go
```

- Você poderá abrir o browser de sua preferência e acessar *localhost:8080*

### Endpoints disponíveis

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