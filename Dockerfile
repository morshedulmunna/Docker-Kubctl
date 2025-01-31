FROM golang:1.20
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o pxomart-api
CMD [""]

EXPOSE 9000