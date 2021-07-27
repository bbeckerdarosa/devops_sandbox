# Deploy Microservice in Golang on GCP

## Instructions

- You should to have install *SDK gcloud* and *Packer*

### App Engine

- Inside folder *app-engine*, to open the terminal and type the following commands below:

```
    gcloud init
```

```
    gcloud app deploy app.yaml
```

- The service will be available in:

```
    {your_instance_URL}/calc/{operation}/{firstNumber}/{secondNumber}
```

and

```
    {your_instance_URL}/calc/history
```

- The values of *operation* are: ```sum```, ```sub```, ```mult``` and ```div```

### Compute Engine with Load Balancer

- To build the GCP image with Packer, you should to do the download of your GCP credential in JSON format. To go ```API's and Services```, ```credentials``` and to create one new key ```Service account key```.

- Inside principal project folder, to open the terminal and type the following command:

```
    packer build -var 'project_id={your_id_project}' -var 'account_file={your_key_file}.json' -var 'ssh_username={your_username}' bake-microservice.json
```

- The packer image will to start built.
- To go in ```Network```, ```Firewall rules``` and to create a rule enabling the port 8080 with TCP protocol to ensure only this port will be accessed.

- Now, you should to go in ```Instance template``` and to create a instance model from GCP image to create with Packer.

- To go in ```Instance groups``` and to create a instance group managed to allow to operate an application on identical Virtual Machine.

- To create a ```Load Balancing``` TCP to do the automatic balancing to instance groups.

- The microservice will be available in: 

```
    {your_instance_URL_external}:8080/calc/{operation}/{firstNumber}/{secondNumber}
```

- The reference values of *operation* are: ```sum```, ```sub```, ```mult``` and ```div```

- To check the history of microservice, access:

```
    {your_instance_URL_external}:8080/calc/history
```

## Presentation

[Video](https://drive.google.com/file/d/1Q76TMeeir_f8OWz5QEWdnC4saM9A_s4e/view)