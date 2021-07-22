# Microservice Calculator in Golang using Vagrant and Ansible

## Instructions

- You need install the Vagrant.
- With the Vagrant installed, access the project directory.
- To start the service, start the Vagrant using the following command:

 ```
    vagrant up
 ```

- The service will be available in:

```
    http://55.55.55.5:8080/
```

- If you wish to stop the service, type the following command:

 ```
    vagrant halt
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