FROM golang:1.22 as build
LABEL authors="Danila Bogdanov https://t.me/DBugdanov"

WORKDIR /go/src/app
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o /go/bin/app cmd/main.go

FROM gcr.io/distroless/base-nossl-debian12
COPY --from=build /go/bin/app .

EXPOSE 8080
ENV PORT=8080
ENV JWT_SECRET="addHereYourJWTSecret"

ENTRYPOINT ["/app"]