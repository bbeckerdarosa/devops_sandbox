# Descomplicando Terraform - Turma #ChamaAsMina
## Instruções usando provider AWS
- Antes de iniciar o terraform, você deverá exportar sua access e secret key para acesso a sua conta aws:
```
export AWS_ACCESS_KEY_ID="sua_access_key"
```
```
export AWS_SECRET_ACCESS_KEY="sua_secret_key"
```

- Com isso, você conseguirá iniciar o terraform:
```
terraform init
```