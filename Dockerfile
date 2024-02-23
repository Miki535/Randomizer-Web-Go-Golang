FROM golang:latest

COPY . /Randomizer-Web-Go-Golang


WORKDIR /Randomizer-Web-Go-Golang

RUN go build -o main main.go

CMD ["./main"]

EXPOSE 8001

