FROM golang:latest

# Add agent requirements
ADD . /go/src/unikernel_agent
WORKDIR /go/src/unikernel_agent
COPY *.go ./
RUN go mod download
RUN go build -o ./main

# Add java runtime building requirements
EXPOSE 6000
CMD ["./main"]
