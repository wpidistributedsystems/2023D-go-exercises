FROM golang:latest

WORKDIR ./src

COPY . .

RUN go mod download && go mod verify

# Edit this line to reflect the path to the go files you want to run, then build and run the container
CMD go run ./'Concurrency Channels and WaitGroups Oh my'/*.go