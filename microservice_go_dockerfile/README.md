# Microservice Calculator in Golang with Dockerfile

## Instruções

- Acesse o diretório do projeto e digite o comando:

``` 
    docker build -t calculator .
```

- Para conseguir rodar o script bash, digite o comando:

``` 
    chmod a+x script_run_microservice.sh
``` 

- Este comando dará a permissão necessária para rodar o script bash

- Agora você já pode rodar o comando a seguir e a calculadora em Golang estará disponível:

``` 
    ./script_run_microservice.sh
```

- Ao rodar, você terá 3 opções:

*1 - Iniciar a calculadora*

*2 - Parar a calculadora*

*3 - Verificar o status da calculadora (RUNNING | NOT RUNNING)*

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
