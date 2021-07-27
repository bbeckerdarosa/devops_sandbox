# Microservice using Vagrant

## Instructions

- You need to install the Vagrant.
- With the Vagrant installed, access the project directory.
- To start the VM, type the following command:

```
    vagrant up
```

- To enter in VM type the following command:

```
    vagrant ssh
```

- To run the microservice in Golang inside the VM, type the following command:

```
    ./script_run_microservice.sh
```

- To run, you will be 3 options:

*1 - Start the calculator*

*2 - Stop the calculator*

*3 - To check the status of the calculator (RUNNING | NOT RUNNING)*


- To go out inside VM, type in terminal:

```
    logout
```

- If you want to delete the VM, just type:

```
    vagrant destroy
```

### Available Endpoints

#### Sum

```
    http://55.55.55.5:8080/calc/sum/{firstNumber}/{secondNumber}
```

#### Sub

```
    http://55.55.55.5:8080/calc/sub/{firstNumber}/{secondNumber}
```

#### Div

```
    http://55.55.55.5:8080/calc/div/{firstNumber}/{secondNumber}
```

#### Mult

```
    http://55.55.55.5:8080/calc/mult/{firstNumber}/{secondNumber}
```

#### History

```
    http://55.55.55.5:8080/calc/history
```

- *firstNumber* matches the first number to be calculated

- *secondNumber* matches the second number to be calculated

