# tart by building the application.
# Build em gocorreio.frete com distroless
FROM golang:1.14.1 as builder

WORKDIR /go/src/gocorreio.frete

COPY gocorreio.frete .

ENV GO111MODULE=on

#RUN go install -v ./...
#RUN GOOS=linux go  build -ldflags="-s -w" -o gocorreio.frete main.go
RUN cp gocorreio.frete /go/bin/gocorreio.frete

RUN ls -lh

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/gocorreio.frete /
CMD ["/gocorreio.frete"]