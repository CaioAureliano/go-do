FROM golang:1.17.8-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /godo cmd/main.go

CMD ["/godo"]