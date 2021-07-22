# Microservice Calculator in Golang on AWS

### Instructions

#### Configuring image AWS of microservice using Packer and Ansible

- You need to have the Packer and Terraform package installed.
- Access the principal directory of project, open the terminal and type the following command to do the creation of image AWS with Packer:

```
    packer build -var 'access_key=your_acess_key' -var 'secret_key=your_secret_key' bake-microservice.json
```

#### Configuring ELB, SG, LC and ASG with Terraform from image AWS created with Packer.

- Open the file *variables.tf* and type the correct AMI ID in *aws_amis*
- In AWS, go to *NETWORK & SECURITY* and  click in *Key Pairs* 
- Create one new Key Pair to get to access the machines with SSH
- Done that, in the terminal type the following command:

```
    terraform apply
```

- You should to put your access keys of AWS, the name of your file .pem created and your IP in this format:

```["{your_ip_range}"]```

- Will be create two machines from image AWS made with Packer.

### Using microservice of AWS from Load Balancer

- To access the calculator, access in browser:

*{your_DNS_name}:8080/calc/{operation}/{firstNumber}/{secondNumber}* 

- The available operations are: sum | sub | mult | div

- You can to see history of the operations in: 

*{your_DNS_name}:8080/calc/history*

### Presentation showing how to run the microservice in AWS:

[Video](https://drive.google.com/file/d/1Y9mqE5bkojkbAKLDtCZsS4Focn95e4Ml/view)

