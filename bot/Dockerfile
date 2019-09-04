FROM golang:1.13.0-buster
RUN mkdir /bot
ADD . /bot/
WORKDIR /bot
RUN go mod download
RUN go build -o main
CMD ["/bot/main"]