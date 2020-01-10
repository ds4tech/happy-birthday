# happy-birthday

## Introduction <a name="intro"></a>

Description: Saves/updates the given user's name and date of birth in the database
  Request: PUT /hello/Morty { "dateOfBirth": "2000-01-01" }
  Response: 204 No Content

Description: Return a hello/birthday message for the given user
  Request: GET /hello/Morty
  Response: 200 OK

a. when Morty’s birthday is in 5 days:
{ "message": "Hello, Morty! Your birthday is in 5 days" }
b. when Morty's birthday is today:
{ "message": "Hello, Morty! Happy birthday" }

## BUILD <a name="build"></a>

### Executable <a name="build.exe"></a>
```
go build -o main cmd/happybirthday/main.go
./main

http://localhost:8888/
```

### Docker container <a name="build.docker"></a>
```
docker build -t happy-birthday -f build/package/Dockerfile .
docker run -d -p 80:8888 happy-birthday
docker tag happy-birthday:latest happy-birthday:latest

http://localhost/
```
