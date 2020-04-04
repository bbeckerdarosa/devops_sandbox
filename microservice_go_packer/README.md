# Packer config file for backing a docker image of microservice Calculator in Golang

## Instruções

- Você precisa instalar o Packer.
- Como Packer instalado, acesse o diretório do projeto.
- Para fazer o build da imagem docker configurada com o Packer, digite o comando:

```
    packer build template-packer.json
```

- Para rodar a imagem docker, digite:

```
    docker run -p 8080:8080 barbbecker/calculator-go:0.1
```

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

- *firstNumber* corresponde ao primeiro número a ser calculado

- *secondNumber* corresponde ao segundo número a ser calculado