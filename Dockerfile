FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

RUN go mod download

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

COPY . .

WORKDIR /app/cmd/ordersystem

EXPOSE 8080
EXPOSE 8000
EXPOSE 50051
