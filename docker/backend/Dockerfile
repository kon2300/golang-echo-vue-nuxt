FROM golang:1.18-rc-alpine

WORKDIR /go/src/backend

RUN apk upgrade --update && \
    apk --no-cache add git

COPY ./backend /go/src/backend

RUN go install github.com/labstack/echo/v4

# CMD ["air"]