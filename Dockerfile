FROM golang
WORKDIR /go
COPY moj.go /go
RUN apt-get update && apt-get install -y nginx
COPY default /etc/nginx/sites-enabled
COPY run.sh /go/run.sh
RUN chmod +x /go/run.sh
RUN go get github.com/gocql/gocql
CMD [ "/go/run.sh" ]
