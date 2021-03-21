FROM golang:1.15-alpine AS builder
RUN apk add git
RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go get -u github.com/shawn-ogg/machineid
RUN go build -o server .

FROM alpine

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app/server .

CMD [ "/app/server" ]
