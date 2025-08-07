FROM golang:1.20

WORKDIR /app
COPY . .
RUN go mod init blog-service && go mod tidy
RUN go build -o blog-service
CMD ["./blog-service"]