FROM golang:1.22.4-alpine as builder

WORKDIR /opt/powclient

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o ./bin/powclient ./cmd/client


FROM golang:1.22.4-alpine

WORKDIR /opt/powclient

ENV WOWPOW_PORT=15100
ENV WOWPOW_HOST=host.docker.internal
ENV WOWPOW_REQUESTS=10

RUN apk add --update \
    bash \
    curl \
    && rm -rf /var/cache/apk/*

COPY --from=builder /opt/powclient/bin/powclient /opt/powclient/bin/powclient

CMD ["/opt/powclient/bin/powclient"]