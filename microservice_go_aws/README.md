# Microservice calculator go on AWS

### Instruções

#### Configurando uma imagem AWS do microsserviço com Packer e Ansible

- Você precisa ter o Packer e o Terraform instalado.
- Acesse o diretório principal do projeto, abra o terminal e digite o seguinte comando para fazer a criação da imagem AWS com Packer:

```
    packer build -var 'access_key=your_acess_key' -var 'secret_key=your_secret_key' bake-microservice.json
```

#### Configurando ELB, SG, LC e ASG com Terraform a partir da Imagem AWS criada com Packer.

- Abra o arquivo *variables.tf* e digite o AMI ID correto em *aws_amis*
- Na AWS, vá até *NETWORK & SECURITY* e clique em *Key Pairs* 
- Crie uma nova Key Pair para conseguir acessar as máquinas com SSH
- Feito isso, no terminal digite o seguinte comando:

```
    terraform apply
```

- Você deverá colocar as suas chaves de acesso da AWS, o nome do seu arquivo .pem criado e o seu IP neste formato:
```["{your_ip_range}"]```

- Será criada duas máquinas a partir da imagem AWS feita com Packer

### Usando o Microsserviço da AWS pelo Load Balancer

- Para acessar a calculadora acesse:

*{your_DNS_name}:8080/calc/{operation}/{firstNumber}/{secondNumber}* 

- As operações disponíveis são: sum | sub | mult | div

- Você também pode ver o histórico das operações realizadas em:

*{your_DNS_name}:8080/calc/history*

### Apresentação mostrando como rodar o microsserviço na AWS

[Vídeo](https://drive.google.com/file/d/1Y9mqE5bkojkbAKLDtCZsS4Focn95e4Ml/view)

