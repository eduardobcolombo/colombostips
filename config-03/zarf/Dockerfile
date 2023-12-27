ARG go=1.20
FROM golang:$go-alpine AS builder

ENV CGO_ENABLED=0

WORKDIR /src
COPY . .

RUN go version
RUN go env

RUN go build main.go

FROM scratch AS config

COPY --from=builder /src/config /config

ENTRYPOINT ["/config"]
