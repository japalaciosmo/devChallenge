FROM golang:1.13.0-buster
RUN mkdir /api
ADD . /api/
WORKDIR /api
RUN go mod download
RUN go build -o main
CMD ["/api/main"]