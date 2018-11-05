FROM golang:1.10-alpine3.8 as builder

RUN apk add git

WORKDIR $GOPATH/src/github.com/agrrh/dummy-service-golang

COPY . ./
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o daemon .
RUN mv daemon /usr/bin/daemon

FROM scratch
COPY --from=builder /usr/bin/daemon ./
CMD ["./daemon"]
