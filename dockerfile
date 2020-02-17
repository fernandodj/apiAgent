# Start from the latest golang base image
FROM golang:1.13

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/agentApi

COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

EXPOSE 80

RUN cd api && go build -o ../main

CMD ./main
