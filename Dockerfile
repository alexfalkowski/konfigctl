FROM golang:1.23.2-bullseye AS build

ARG version=latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w -X 'github.com/alexfalkowski/konfigctl/cmd.Version=${version}'" -a -o konfigctl main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=build /app/konfigctl /konfigctl
ENTRYPOINT ["/konfigctl"]
