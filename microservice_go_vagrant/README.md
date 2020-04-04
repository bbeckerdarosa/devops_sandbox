# Microservice Calculator in Golang with Vagrant

## Instruções

- Você precisa instalar o Vagrant.
- Com o Vagrant instalado, acesse o diretório do projeto.
- Para iniciar a VM digite no terminal:

```
    vagrant up
```

- Para entrar na VM digite:

```
    vagrant ssh
```

- Para rodar o microservice em Golang dentro da VM, digite:

```
    ./script_run_microservice.sh
```

- Ao rodar, você terá 3 opções:

*1 - Iniciar a calculadora*

*2 - Parar a calculadora*

*3 - Verificar o status da calculadora (RUNNING | NOT RUNNING)*


- Para sair da VM, digite:

```
    logout
```

- Caso você queira excluir a VM, basta digitar:

```
    vagrant destroy
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

