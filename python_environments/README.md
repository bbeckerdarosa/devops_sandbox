## Python

### Instructions

- To run the following commands, you need to install *python*

#### List environment variables

- Enter the command at the terminal:

```
    python env_service1.py
```

- Open the browser and put in the url:

```
    http://localhost:8080/conf/env
```

#### Create environment variable

- Enter the command in the terminal:

```
    python create_service2.py
```

- Open postman or browser and put in the url:

```
    http://localhost:8080/env/{env_nome}/{env_var}
```

- POST method
- Replace the names *env_nome* and *env_var* to create your variable.

#### List running software (PID and cmd) in a Slack channel

- You must first create a new channel in Slack.
- Configure the Incoming WebHooks app on *Channel Settings* and *Add apps...*
- Choose a channel to apply the WebHook addition and click *Add Incoming WebHokks integration*
- Copy the webhook url that will appear.
- At the script *slack_service3.py* paste the webhook url into *SLACK_DEBUG_URL*
- Put the name of the channel you want to use in *SLACK_DEBUG_CHANNEL*

- To run the python script, type the following command at the terminal:

```
    python slack_service3.py
```

- Enter your channel via slack and you will see the message with running software.