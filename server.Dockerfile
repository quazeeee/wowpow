FROM golang:1.22.4-alpine as builder

WORKDIR /opt/powservice

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o ./bin/powservice ./cmd/server


FROM golang:1.22.4-alpine

WORKDIR /opt/powservice

ENV WOWPOW_PORT=15100
EXPOSE 3000

RUN apk add --update \
    bash \
    curl \
    && rm -rf /var/cache/apk/*

COPY --from=builder /opt/powservice/bin/powservice /opt/powservice/bin/powservice

CMD ["/opt/powservice/bin/powservice"]