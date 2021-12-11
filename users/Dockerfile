FROM golang:1.16-alpine AS builder

ENV GO111MODULE=on
ENV GOFLAGS="-mod=vendor"

WORKDIR /go/src/service

COPY . .

WORKDIR /go/src/service/cmd/app
RUN GOOS=linux CGO_ENABLED=0 go build -o service

WORKDIR /go/src/service

FROM scratch

FROM builder AS dev

COPY --from=builder /go/src/service/cmd/app/service /go/bin/service
COPY --from=builder /go/src/service/config/development.yaml /go/bin/config/development.yaml

EXPOSE 8080

ENTRYPOINT ["/go/bin/service"]