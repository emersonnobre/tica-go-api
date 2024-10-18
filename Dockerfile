FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest 

RUN make run-swagger

RUN go build -o main ./src/main.go
EXPOSE 3000

CMD ["./main"]