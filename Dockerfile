FROM golang
WORKDIR /go
COPY moj.go /go
COPY run.sh /go
RUN go get github.com/gocql/gocql
CMD [ "go", "run", "/go/moj.go" ]
