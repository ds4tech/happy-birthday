# happy-birthday

1. [Introduction](#intro)
2. [Build](#build) <br>
   2.1. [Exec](#build.exe) <br>
   2.2. [Docker](#build.docker)
3. [Deploy](#deploy) <br>
 3.1. [Kubernetes](#deploy.k8s) <br>
 3.2. [AWS ECS](#deploy.ecs)
4. [Usage](#usage)
5. [Continous Integration](#ci)


## Introduction <a name="intro"></a>

Description: Saves/updates the given user's name and date of birth in the database <br>
  Request: PUT /hello/Morty { "dateOfBirth": "2000-01-01" } <br>
  Response: 204 No Content <br>

Description: Return a hello/birthday message for the given user <br>
  Request: GET /hello/Morty <br>
  Response: 200 OK <br>

a. when Morty’s birthday is in 5 days:<br>
```
{ "message": "Hello, Morty! Your birthday is in 5 days" }
```
b. when Morty's birthday is today: <br>
```
{ "message": "Hello, Morty! Happy birthday" }
```

## BUILD <a name="build"></a>

### Executable <a name="build.exe"></a>
```
go build -o main cmd/happybirthday/main.go
./main

curl http://localhost:8888/
```

### Docker container <a name="build.docker"></a>
```
docker build -t happy-birthday -f build/package/Dockerfile .
docker run -d -p 80:8888 happy-birthday
docker tag happy-birthday:latest happy-birthday:latest

curl http://localhost/
```

## DEPLOY <a name="deploy"></a>

### Kubernetes <a name="deploy.k8s"></a>
```
kubectl create -f deployments/kubernetes/k8s-replicaSet.yml

curl http://192.168.99.100:30000/
```

### AWS ECS <a name="deploy.ecs"></a>
```
cd deployments/aws/ecs
terraform init
terraform apply
//copy output: repository url
cd ../../../
$(aws ecr get-login --no-include-email --region eu-west-1)
docker build -t happy-birthday .
docker tag happy-birthday:latest 311744426619.dkr.ecr.eu-west-1.amazonaws.com/happy-birthday:latest
docker push 311744426619.dkr.ecr.eu-west-1.amazonaws.com/happy-birthday:latest
```

## USEAGE <a name="usage"></a>


## Continous Integration <a name="ci"></a>
Pipeline script written in Groovy is placed in [build/ci directory](https://github.com/ds4tech/happy-birthday/blob/master/build/ci/pipeline.yaml). It is dedicated for Jenkins Pipeline JOB. <br>
To run Jenkins on AWS, run terraform scripts in ```deployments/aws/jenkins-ec2_instance```
