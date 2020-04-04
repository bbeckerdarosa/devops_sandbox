# Jenkins Job - Deploy microservice Go on Kubernates

## Instalação

- [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)
- [Jenkins Kubernates - plugin](https://github.com/jenkinsci/kubernetes-plugin)
- [Jenkins Kubernates CLI - plugin](https://github.com/jenkinsci/kubernetes-cli-plugin/blob/master/README.md)

## Instruções para deploy no Jenkins

- Na pasta onde você colocou o .war do Jenkins, rode o comando:

```
    java -jar jenkins.war --httpPort=8180
```

- O Jenkins estará disponível em *http://localhost:8180/*

### Configurações

- Crie uma credencial do tipo `certificate` configurada com o nome `kubernetes-certificate`
    - Use como password: `secret`
    - Faça o upload do certificado que se encontra no diretório: `~/.minikube/minikube.pfx`
    - Vá em `Manage Jenkins` > `Configure System`
    - Bem embaixo, em `Cloud `, selecione `Add a new cloud`
    - Em `Kubernetes server certificate key`, você deverá colocar a chave do certificado que pode ser encontrada em `~/.minikube/ca.crt`

### Pipeline Jenkins

- Clique na opção `GitHub Project` e insira a url do repositório do github *https://github.com/barbbecker/devops_sandbox.git*

- Em `Pipeline`, selecione a opção `Pipeline script from SCM`
    - Em `SCM` selecione a opção `Git`
    - Clique em `Apply` e `Save`

- Clique em `Build with Parameters` e coloque a url: `https://192.168.99.100:8443`

- A aplicação estará disponível acessando a url: `https://192.168.99.100:<service_port>`
