FROM golang:latest
RUN apt-get update -y && apt-get install sqlite3 libsqlite3-dev -y
RUN mkdir /app
WORKDIR /app
ADD . /app
VOLUME /app
RUN go mod download
EXPOSE 8080
CMD "go" "run" "server.go"