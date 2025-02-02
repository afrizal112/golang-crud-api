FROM golang:1.16-alpine

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main .

EXPOSE 8081

CMD ["./main"]