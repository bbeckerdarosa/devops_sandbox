# Microservice Calculator in Golang using Docker/Dockerfile

## Instruções

- Access the project directory and type the following command:

``` 
    docker build -t calculator .
```

- To run the bash script, type in the terminal:

``` 
    chmod a+x script_run_microservice.sh
``` 

- This command give the permission to run the bash script:

- Now, you can to run the following command and the golang calculator will be available:

``` 
    ./script_run_microservice.sh
```

- To run, you will 3 options:

*1 - Start the calculator*

*2 - Stop the calculator*

*3 - To check the calculator status (RUNNING | NOT RUNNING)*

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

- *firstNumber* matches the first number to be calculated

- *secondNumber* matches the second number to be calculated