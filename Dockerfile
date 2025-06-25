FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go build -o panic-app .

EXPOSE 8080

ENTRYPOINT ["./panic-app"]
