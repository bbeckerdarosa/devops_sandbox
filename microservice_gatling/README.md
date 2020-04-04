# Gatling Stress and Chaos Test

## Instruções

- Para buildar o projeto você precisa ter o [SBT](https://www.scala-sbt.org/) instalado

## Buildando o projeto localmente

- Faça o clone do projeto privado em:

```
    git clone https://github.com/barbbecker/devops_sandbox.git
```

- Para rodar a calculadora digite:

```
    go run calculator.go
```

- Para rodar o Gatling test definindo a quantidade de usuários e o tempo, digite o comando:

```
    sbt -Dusers=${params.USERS} -Dduration=${params.DURATION} gatling:test
```

- Substitua `params.USERS` pela quantidade de usuários e `params.DURATION` pelo tempo em segundos que desejar

- Os resultados estarão disponíveis, acessando os seguintes diretórios:

`target/gatling/calculatorservicestresstest-{timestamp}/simulation.log`

## Buildando o projeto com o Jenkins

- Você também pode buildar o projeto com o Jenkins, usando o arquivo `Jenkinsfile`