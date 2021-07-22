# Gatling Stress and Chaos Test

## Instructions

- To build the project you need to have [SBT](https://www.scala-sbt.org/) installed

## To build the project locally

- Clone the public project in:

```
    git clone https://github.com/bbeckerdarosa/devops_sandbox.git
```

- To run the calculator, type in the terminal:

```
    go run calculator.go
```

- To run the Gatling test defining the quantity of users and the time, type in the terminal:

```
    sbt -Dusers=${params.USERS} -Dduration=${params.DURATION} gatling:test
```

- Replace `params.USERS` with quantity of users and `params.DURATION` with time in seconds

- The results will be available accessing the following directories:

`target/gatling/calculatorservicestresstest-{timestamp}/simulation.log`

## To build the project using Jenkins

- You can to build the project with Jenkins, using the file `Jenkinsfile`