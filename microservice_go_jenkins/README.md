# Pipeline Jenkins to bake image with Packer and deploy Microservice in Golang

### Tools to need to be installed first:

- [Packer](https://www.packer.io/)
- [Docker](https://www.docker.com/)
- [Jenkins](https://jenkins.io/)

### Instructions

- In folder where you to put the jenkins .war type the following command: 

```
    java -jar jenkins.war --httpPort=8180
```

- The Jenkins will be available in *http://localhost:8180/*

- At the Jenkins, you should to create 2 pipelines:

#### Bake
- Responsable to prepare an Docker image provisioned by Packer.

- To click at option ```Github Project``` and insert the url of private repository github *https://github.com/bbeckerdarosa/devops_sandbox.git*
- To copy the script of Bake/Jenkinsfile and paste at option "Pipeline script".
- At the preparetion stage, the Jenkins will make the clone Github repository.
- At the Bake stage, the Jenkins will make the image bake Docker with Packer.

#### Launch
- Responsable to do the microservice deploy in Golang.

- To copy the script Launch/Jenkinsfile and paste at the option "Pipeline script".
- At the Launch stage, the Jenkins will make the microservice deploy in Golang with image name to create in Bake step.

- The microservice will be available in *http://localhost:8080/*

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