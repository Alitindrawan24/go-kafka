FROM golang:1.19-alpine AS builder

ARG TARGETARCH
RUN echo "$TARGETARCH"

RUN apk update && apk add --no-cache \
    gcc \
    musl-dev \
    # explicitly install SASL package
    cyrus-sasl-dev

WORKDIR /build
COPY . .

RUN go mod vendor

RUN CGO_ENABLED=1 \
    GO111MODULE=on \
    GOOS=linux \
    GOARCH="$TARGETARCH" \
    # explicitly link to libsasl2 installed as part of cyrus-sasl-dev
    CGO_LDFLAGS="-lsasl2" \
    go build -mod=vendor -o /build/main -tags musl -ldflags "-w -s" .

CMD ["./main"]