FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build calculator-service .
CMD ["/app/calculator-service"]