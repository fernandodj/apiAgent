# Start from the latest golang base image
FROM golang:1.13

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/apiAgent

COPY . .

RUN mkdir compliance_result
# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

EXPOSE 80

RUN cd app && go build -o ../main

CMD ./main




