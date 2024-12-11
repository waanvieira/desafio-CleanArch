FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

RUN go mod download

COPY . .

WORKDIR /app/cmd/ordersystem

EXPOSE 8080
EXPOSE 8000
EXPOSE 50051

# Command executed when we create the build, it a dummy command
CMD ["tail", "-f", "/dev/null"]
