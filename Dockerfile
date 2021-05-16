FROM golang:latest
RUN apt-get update && apt-get upgrade -y
RUN apt-get install -y curl gnupg2
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz | tar xvz
RUN mv ./migrate.linux-amd64 /usr/bin/migrate

RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app
VOLUME /go/src/app/
RUN go mod download
EXPOSE 8080
CMD "go" "run" "main.go"