## Python

### Instruções

- Para rodar os comandos a seguir, você precisa instalar o *python*

#### Listar variáveis de ambiente

- Digite pelo terminal o comando:

```
    python env_service1.py
```

- Abra o browser e coloque na url:

```
    http://localhost:8080/conf/env
```

#### Criar variável de ambiente

- Digite pelo terminal o comando:

```
    python create_service2.py
```

- Abra o postman ou browser e coloque na url:

```
    http://localhost:8080/env/{env_nome}/{env_var}
```

- Método POST
- Substitua os nomes *env_nome* e *env_var* para criar sua variável.

#### Listar os softwares em execução (PID e cmd) em um channel do Slack

- Você deve primeiro criar um novo channel no Slack.
- Configure o app Incoming WebHooks em *Channel Settings* e *Add apps...*
- Escolha um channel para aplicar a adição do WebHook e clique em *Add Incoming WebHokks integration*
- Copie a url do Webhook que vai aparecer
- No script *slack_service3.py* cole a url do Webhook em *SLACK_DEBUG_URL*
- Coloque o nome do channel que você deseja usar em *SLACK_DEBUG_CHANNEL*

- Para rodar o script python, digite o comando a seguir pelo terminal:

```
    python slack_service3.py
```

- Entre no seu channel pelo slack e você verá a mensagem com os softwares em execução