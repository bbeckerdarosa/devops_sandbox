#!/bin/bash

echo "Welcome to the calculator!"

echo "Type it the first number:"
read first
echo "Type it the second number:"
read second

addition() {
    echo "The result of addition "$first" and "$second" is:"
    result=$(($first + $second))
    echo $result
}

subtraction() {
    echo "The result of subtraction "$first" and "$second" is:"
    result=$(($first - $second))
    echo $result
}

multiplication() {
    echo "The result of multiplication "$first" and "$second" is:"
    result=$(($first * $second))
    echo $result
}

division() {
    echo "The result of division " $first " and " $second " is:"
    result=$(($first / $second))
    echo $result
}


while true;
do 
echo "Choose the operation:"
echo "1 - addition"
echo "2 - subtraction"
echo "3 - multiplication"
echo "4 - division"
read operation 

case $operation in
    1) 
        addition
        ;;
    2)
        subtraction
        ;;
    3) 
        multiplication
        ;;
    4) 
        division
        ;;
    *)
        echo "Try again!"
esac
done

