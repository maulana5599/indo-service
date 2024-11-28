from golang:alpine as builder

WORKDIR /app
ADD . .
RUN go build -o /usr/local/bin/myapp

EXPOSE 7004
CMD ["/usr/local/bin/myapp"]
