FROM golang:1.22 as build
LABEL authors="Danila Bogdanov https://t.me/DBugdanov"

WORKDIR /go/src/app
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download
COPY ../* ./
RUN CGO_ENABLED=1 GOOS=linux go build -o /go/bin/app
EXPOSE 8080

FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/app /
ENV PORT 8080
CMD ["/app"]
