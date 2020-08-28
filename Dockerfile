FROM golang:1.12-alpine as builder

WORKDIR /go/src/app

COPY . .

RUN go get -v ./...  \
    && go build

FROM scratch

COPY --from=builder /go/src/app/app .

ENTRYPOINT ["./app"]

