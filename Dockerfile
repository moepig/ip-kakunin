FROM golang:1.21-buster AS build

WORKDIR /build

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /app

# ---------------
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app /app

USER nonroot:nonroot

ENTRYPOINT ["/app"]