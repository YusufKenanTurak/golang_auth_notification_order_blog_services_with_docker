FROM golang:1.20

WORKDIR /app
COPY . .
RUN go mod init auth-service && go mod tidy
RUN go build -o auth-service
CMD ["./auth-service"]