# NoSQL 

- Roles Ansible customizados para provisionamento

## ROLES

- Os seguintes arquivos podem ser encontrados no diretório `roles`

### base
    - Update do Sistema Operacional

### jdk
    - Responsável por instalar o JDK 8, para trabalhar com o Kafka e Cassandra

### redis
    - Responsável por instalar o redis e startar o serviço do redis

### elasticsearch
    - Responsável por instalar o apt-transport-https, elasticsearch e startar o serviço do elasticsearch
  
### cass
    - Responsável por instalar o apt-transport-https, cassandra e startar o serviço do Cassandra

### kafka
    - Responsável por instalar o Zookeeper, extrair o Kafka para o caminho /opt/kafka/ e startar o serviço do Kafka

## Instruções para rodar em um container Docker

- Você deverá acessar o diretório onde está seu arquivo `playbook.yml` do Ansible

- Abra o terminal e digite o comando:

```
    docker build -t barbbecker/tema17 .
```

- Este comando irá buildar a imagem com o Dockerfile
- Em seguida, digite o comando:

```
    sudo docker run -v "$(pwd)":/tmp --workdir="/tmp" -it barbbecker/tema17
```

- Este comando passará seus arquivos para dentro do container e executará as tasks de cada /role do ansible que está no `playbook.yml`

*Utilizei este repositório como exemplo: `https://github.com/William-Yeh/docker-ansible`*
