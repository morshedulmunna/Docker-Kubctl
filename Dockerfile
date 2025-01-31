FROM golang:1.20
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o user-service
CMD ["./user-service"]

EXPOSE 9000