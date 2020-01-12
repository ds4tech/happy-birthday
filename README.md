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
It is required to first build the image and push it to any container repository. <br>
Second step is to update yaml file k8s-replicaSet.yml in line 20 and change repository name from ds4tech to the one where docker image has been uploaded.

```
kubectl create -f deployments/kubernetes/k8s-replicaSet.yml

curl http://192.168.99.100:30000/
```

### AWS ECS <a name="deploy.ecs"></a>
```
$(aws ecr get-login --no-include-email --region eu-west-1)
cd deployments/aws/ecs-jenkins
terraform init
terraform apply
```
Copy output ip and log in to the machine using generated key:
```
ssh -i mykey ubuntu@34.244.38.111
sudo su -
cat /var/lib/jenkins/secrets/initialAdminPassword
```
Copu the password and pass it to web browser

```
su - jenkins
aws configure

AWS Access Key ID [None]: ***
AWS Secret Access Key [None]: ***
Default region name [None]: eu-west-1
```
Jenkins is now set up and ready to run jobs.


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

## Continous Integration <a name="ci"></a>
Pipeline script written in Groovy is placed in [build/ci directory](https://github.com/ds4tech/happy-birthday/blob/master/build/ci/pipeline.yaml). It is dedicated for Jenkins Pipeline JOB. <br>
