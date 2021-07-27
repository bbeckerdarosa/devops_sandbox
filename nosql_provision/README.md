# Provision NoSQL using Ansible and Docker 

- Roles Ansible customized to provisioning

## ROLES

- The following files can be found at the directory `roles`

### base
    - Operating System Update

### jdk
    - Responsible for installing JDK 8, to work with Kafka and Cassandra

### redis
    - Responsible for installing redis and starting the redis service

### elasticsearch
    - Responsible for installing apt-transport-https, elasticsearch and starting the elasticsearch service
  
### cass
    - Responsible for installing apt-transport-https, cassandra and starting the Cassandra service

### kafka
    - Responsible for installing Zookeeper, extracting Kafka to /opt/kafka/ path and starting Kafka service

## Instructions for running in a Docker container

- You should access the directory where your file is `playbook.yml`

- Open the terminal and type the command:

```
    docker build -t barbbecker/tema17 .
```

- This command will build the image with the Dockerfile
- Then enter the command:

```
    sudo docker run -v "$(pwd)":/tmp --workdir="/tmp" -it barbbecker/tema17
```

- This command will pass your files into the container and to run the tasks of each ansible /role that is in the `playbook.yml`

*I used this repository as an example: `https://github.com/William-Yeh/docker-ansible`*
