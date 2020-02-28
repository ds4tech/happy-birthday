# happy-birthday
test

1. [Introduction](#intro)
2. [Usage](#usage)
3. [Build](#build) <br>
   3.1. [Exec](#build.exe) <br>
   3.2. [Docker](#build.docker)
4. [Unit Tests](#unittest) <br>
5. [Deploy](#deploy) <br>
 5.1. [Kubernetes](#deploy.k8s) <br>
 5.2. [AWS ECS](#deploy.ecs)
6. [Continous Integration](#ci)
7. [Continous Deployment](#cd) <br>
   7.1. [Manual](#cd.manual) <br>
   7.2. [Automatic](#cd.automatic)


## Introduction <a name="intro"></a>

This application provides a simple microservice which saves user data: name and date of birth. The user data can then be fetched. The "Happy birthday" wishes will be returned depending on today's date. The application is not initialized with any data, so some PUT actions have to be run first.

## USEAGE <a name="usage"></a>

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

## Unit Tests <a name="unittest"></a>

To run unit tests run following command inside folder pkg/birthday.
```
cd pkg/birthday
go test
```

or to measure test coverage running
```
go test -coverprofile=cover.out
go tool cover -html=cover.out
```

## DEPLOY <a name="deploy"></a>

### Kubernetes <a name="deploy.k8s"></a>
It is required to first build the image and push it to any container repository. <br>
Second step is to update yaml file k8s-replicaSet.yml in line 20 and change repository name from ds4tech to the one where docker image has been uploaded.
Once this is done, pods can be created by following command.
```
kubectl create -f deployments/kubernetes/k8s-replicaSet.yml

curl http://192.168.99.100:30000/
```

### AWS ECS <a name="deploy.ecs"></a>
Provide AWS credentials by running following command. Then run Terraform.
```
$(aws ecr get-login --no-include-email --region eu-west-1)
cd deployments/aws/ecs-jenkins
terraform init
terraform apply
```
Copy output ip and log in to the machine using generated key:
```
ssh -i mykey ubuntu@output_ip
sudo su -

su - jenkins
aws configure

AWS Access Key ID [None]: ***
AWS Secret Access Key [None]: ***
Default region name [None]: eu-west-1

cat /var/lib/jenkins/secrets/initialAdminPassword
```
Open web browser and go to http://output_ip:8080/

Copy jenkins initialAdminPassword and pass it to web browser. Click install sugested plugins.

Jenkins is now set up and ready to run jobs.


## Continous Integration <a name="ci"></a>
Pipeline script written in Groovy is placed in [build/ci directory](https://github.com/ds4tech/happy-birthday/blob/master/build/ci/pipeline.yaml). It is dedicated for Jenkins Pipeline JOB. <br>
Update lines 63 and 65 with the repository names.
Then create new pipeline job, name it build-docker-image and paste updated code of pipeline.yaml


## Continous Deployment <a name="cd"></a>
After completion of previous steps the code can be deployed in two ways.

### CD - manual <a name="cd.manual"></a>
Update ```myapp-vars.tf```, set MYAPP_SERVICE_ENABLE=1 and set MYAPP_VERSION=build_number to last succeful build from jenkins.
This will create new revision of task definition, and a new task from it.
Then run terraform again.
```
terraform apply
```

### CD - automatic <a name="cd.automatic"></a>
Create new pipeline job, name it deploy-docker-image and paste code from [aws/ecs-jenkins/deploy/pipeline.yaml](https://github.com/ds4tech/happy-birthday/blob/master/deployments/aws/ecs-jenkins/deploy/pipeline.yaml).
