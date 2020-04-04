# Pipeline Jenkins to bake image with Packer and deploy microservice in Golang

### Ferramentas que precisam ser instaladas antes:

- [Packer](https://www.packer.io/)
- [Docker](https://www.docker.com/)
- [Jenkins](https://jenkins.io/)

### Instruções

- Na pasta onde você colocou o .war do Jenkins, rode o comando: 

```
    java -jar jenkins.war --httpPort=8180
```

- O Jenkins estará disponível em *http://localhost:8180/*

- No Jenkins, você deverá criar 2 pipelines:

#### Bake
- Responsável por preparar uma imagem Docker provisionada pelo Packer.

- Clique na opção GitHub Project e insira a url do repositório privado do github *https://github.com/barbbecker/devops_sandbox.git*
- Copie o script do Bake/Jenkinsfile e cole na opção "Pipeline script".
- No stage de Preparação, o jenkins fará o clone do repositório Github.
- No stage de Bake, o jenkins fará o bake da imagem Docker com Packer.

#### Launch
- Responsável por fazer o deploy do microsserviço em Golang.

- Copie o script do Launch/Jenkinsfile e cole na opção "Pipeline script".
- No stage de Launch, o jenkins fará o deploy do microsserviço em Golang com o nome da imagem criada na etapa de Bake.

- O microsserviço estará disponível em *http://localhost:8080/*

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