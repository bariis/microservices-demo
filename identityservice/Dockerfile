FROM golang:1.16.4-alpine3.13

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 5002

RUN go build -o main .

CMD ./main



