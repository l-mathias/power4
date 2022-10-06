FROM golang:alpine as builder

WORKDIR /build

COPY . .

RUN go build -o power4 grpcConsole/main.go

FROM alpine

COPY --from=builder /build/power4 .

CMD "./power4"