# Microservice Calculator in Golang with Vagrant provisioning Ansible

## Instruções

- Você precisa instalar o Vagrant.
- Com o Vagrant instalado, acesse o diretório do projeto.
- Para iniciar o serviço, inicie o Vagrant como comando:

 ```
    vagrant up
 ```

- O serviço estará disponível em:

```
    http://55.55.55.5:8080/
```

- Caso desejar parar o serviço, digite o comando:

 ```
    vagrant halt
 ```

### Endpoints disponíveis

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

- *firstNumber* corresponde ao primeiro número a ser calculado

- *secondNumber* corresponde ao segundo número a ser calculado