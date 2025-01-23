from golang:alpine as builder

WORKDIR /app
ADD . .
RUN go build -o /usr/local/bin/myapp cmd/main.go

EXPOSE 8090
CMD ["/usr/local/bin/myapp"]
