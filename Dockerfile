FROM golang:1.17-buster AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o go-users ./src

FROM debian:buster-slim

WORKDIR /app
COPY --from=build /app/go-users /app/go-users

CMD ["/app/go-users"]
