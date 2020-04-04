#!/bin/bash

start_microservice(){
  if [ "$(netstat -plnt | grep 8080)" ]; then
    echo "Port 55.55.55.5:8080 is already in use"
  else
    echo "Service running in port 55.55.55.5:8080"
    run_microservice
  fi
}

run_microservice(){
    if [ "$(docker ps -a | grep -c calculator)" == 0 ]; then
        docker run -p 8080:8080 -it --name calculator calculator
        echo "The container calculator has been created and is running!"
    else
        docker start calculator
        echo "Restart microservice calculator!"
    fi
}

stop_microservice(){
    docker stop calculator
    echo "The container calculator was stopped!"
}

status_microservice(){
    if [ "$(docker inspect -f {{.State.Running}} calculator)" == true ]; then
        echo "Microservice is RUNNING!"
    elif [ "$(docker inspect -f {{.State.Running}} calculator)" == false ]; then
        echo "Microservice is NOT RUNNING!"
    else
        echo "Microservice does not exist!"
    fi
}

while true
do
    
echo "Choose the command:"
echo "1 - Start microservice"
echo "2 - Stop microservice"
echo "3 - Check the status of microservice (RUNNING|NOT RUNNING)"

read command 

case $command in
    1)
        start_microservice
        ;;
    2) 
        stop_microservice
        ;;
    3)
        status_microservice
        ;;
    *)
        echo "Try again!"
esac
done
