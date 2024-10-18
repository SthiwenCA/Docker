FROM golang:1.22.1 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

#stage 2
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/myapp .

EXPOSE 8080

CMD [ "./myapp" ]