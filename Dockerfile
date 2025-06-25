FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod init panic-app \
 && go build -o panic-app .

EXPOSE 8080

ENTRYPOINT ["./panic-app"]
