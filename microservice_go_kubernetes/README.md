# Jenkins - Deploy Microservice on Kubernetes

## Installation

- [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)
- [Jenkins Kubernetes - plugin](https://github.com/jenkinsci/kubernetes-plugin)
- [Jenkins Kubernetes CLI - plugin](https://github.com/jenkinsci/kubernetes-cli-plugin/blob/master/README.md)

## Instructions to deploy in Jenkins

- In folder where you to put the jenkins .war type the following command: 

```
    java -jar jenkins.war --httpPort=8180
```

- The Jenkins will be available in *http://localhost:8180/*

### Settings

- To create an credential type `certificate` configured with the name `kubernetes-certificate`
    - To use like password: `secret`
    - To do the certificate upload in: `~/.minikube/minikube.pfx`
    - To go in `Manage Jenkins` > `Configure System`
    - Right down, in `Cloud `, select `Add a new cloud`
    - In `Kubernetes server certificate key`, you should to put the certificate key which can be found in `~/.minikube/ca.crt`

### Pipeline Jenkins

- To click at the option `GitHub Project` and insert the Github repository *https://github.com/bbeckerdarosa/devops_sandbox.git*

- In `Pipeline`, select the option `Pipeline script from SCM`
    - In `SCM` select the option `Git`
    - To click in `Apply` and `Save`

- Click in `Build with Parameters` and to put the url: `https://192.168.99.100:8443`

- The application will be available accessing the url: `https://192.168.99.100:<service_port>`
