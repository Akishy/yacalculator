FROM golang:1.22 as build


WORKDIR /go/src/app
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o /go/bin/app cmd/main.go \
    && mkdir "/database" \
    && chmod 777 "/database"


FROM gcr.io/distroless/base-nossl-debian12:debug
LABEL authors="Danila Bogdanov https://t.me/DBugdanov"
COPY --from=build /go/bin/app .
COPY --from=build /database .

EXPOSE 8081
ENV PORT=8081

ENTRYPOINT ["/app"]