FROM golang:1.18-rc-alpine

WORKDIR /go/src/backend

RUN apk upgrade --update && \
    apk --no-cache add git

COPY ./backend /go/src/backend

RUN go mod tidy && \
    go install github.com/cosmtrek/air@v1.27.3

CMD ["air", "-c", ".air.toml"]