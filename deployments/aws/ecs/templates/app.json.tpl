[
  {
    "essential": true,
    "memory": 256,
    "name": "happy-birthday",
    "cpu": 256,
    "image": "${REPOSITORY_URL}:latest",
    "workingDirectory": "/app",
    "command": ["./main"],
    "portMappings": [
        {
            "containerPort": 8888,
            "hostPort": 8888
        }
    ]
  }
]
