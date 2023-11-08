# tart by building the application.
# Build em gocorreio.frete com distroless
FROM golang:1.20 as builder

WORKDIR /go/src/

COPY . .

RUN pwd
RUN ls -lh
RUN cp gocorreio.frete /go/bin/gocorreio.frete
RUN ls -lh /go/bin
RUN mkdir /go/bin/credentials/
RUN cp credentials/credentials.json /go/bin/credentials/credentials.json
RUN ls -lh /go/bin/credentials

# FROM alpine:latest AS final
# RUN apk update
# RUN apk add --no-cache tzdata
# RUN apk add --no-cache ca-certificates
# ENV TZ="America/Sao_Paulo"
# RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
# COPY --from=builder /go/bin/gocorreio.frete /
# COPY --from=builder /go/bin/credentials/ /
# RUN ls -lh
# RUN ls -lh credentials
# CMD ["/gocorreio.frete"]
