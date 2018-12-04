FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get ./...;go build -o main .
CMD ["/app/main"]