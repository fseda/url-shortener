FROM golang:latest as base

WORKDIR /usr/app

COPY go.mod go.sum ./

COPY ./src/ /usr/app/src
COPY ./static/ /usr/app/static
COPY ./*.go ./

RUN go build -o app

FROM base AS production

EXPOSE 8080

CMD ["./app"]