FROM golang:1.18 as builder

WORKDIR /app


COPY * ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping


EXPOSE 8080

# Run
CMD ["/docker-gs-ping"]