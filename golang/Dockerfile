FROM golang:1.17.1-alpine3.13 as builder

WORKDIR /go/src
COPY ./src/ .

RUN go build main.go

FROM alpine:3.13
COPY --from=builder /go/src/main /main
CMD [ "/main" ]