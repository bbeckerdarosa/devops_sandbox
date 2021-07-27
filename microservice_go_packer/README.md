# Microservice Calculator in Golang using Packer and Docker

## Instructions

- You need to install the Packer.
- With the Packer installed, access the project directory.
- To do the image build docker configured with Packer, type the following command:

```
    packer build template-packer.json
```

- To run the Docker image, type the following command:

```
    docker run -p 8080:8080 barbbecker/calculator-go:0.1
```

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