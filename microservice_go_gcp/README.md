# Deploy your Go microservice on GCP

## Instruções

- Você deverá ter instalado *SDK gcloud* e *Packer*

### App Engine

- Dentro da pasta *app-engine*, abra o terminal e digite os comandos, um depois do outro:

```
    gcloud init
```

```
    gcloud app deploy app.yaml
```

- O serviço estará disponivel em:

```
    {your_instance_URL}/calc/{operation}/{firstNumber}/{secondNumber}
```

e

```
    {your_instance_URL}/calc/history
```

- Os valores referentes a *operation* são: ```sum```, ```sub```, ```mult``` e ```div```

### Compute Engine with Load Balancer

- Para buildar a imagem GCP com o Packer, você deve fazer o download da sua credencial da GCP em formato JSON. Vá até ```API's e Services```, ```credentials``` e crie uma nova chave ```Service account key```.

- Dentro da pasta principal do projeto, abra o terminal e digite o comando:

```
    packer build -var 'project_id={your_id_project}' -var 'account_file={your_key_file}.json' -var 'ssh_username={your_username}' bake-microservice.json
```

- A imagem packer começará a ser construída.
- Vá até ```Network```, ```Firewall rules``` e crie uma regra habilitando a porta 8080 com o protocolo TCP para garantir que somente esta porta será acessada.

- Agora você deverá ir até ```Instance template``` e criar um modelo de instância a partir da imagem GCP criada com o Packer.

- Vá ate ```Instance groups``` e crie um grupo de instâncias gerenciadas que permitem operar uma aplicação em VM's idênticas.

- Crie um ```Load Balancing``` TCP para fazer o balanceamento automático para os grupos de instâncias.

- O microsserviço estará disponível em: 

```
    {your_instance_URL_external}:8080/calc/{operation}/{firstNumber}/{secondNumber}
```

- Os valores referentes a *operation* são: ```sum```, ```sub```, ```mult``` e ```div```

- Para verificar o histórico do microsserviço, acesse:

```
    {your_instance_URL_external}:8080/calc/history
```

## Apresentação

[Video](https://drive.google.com/file/d/1Q76TMeeir_f8OWz5QEWdnC4saM9A_s4e/view)