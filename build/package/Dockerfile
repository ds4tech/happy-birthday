FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Mateusz Szymczyk"

# Set the Current Working Directory inside the container
WORKDIR /app
#WORKDIR $GOPATH/src/github.com/ds4tech/happy-birthday

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download

# Build the Go app
RUN go build -o main cmd/happybirthday/main.go

# Expose port 8888 to the outside world
EXPOSE 8888

# Command to run the executable
CMD ["./main"]
