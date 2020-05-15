FROM golang:alpine
RUN apk add --no-cache git gcc musl-dev
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN CGO_ENABLED=1 go install
ENTRYPOINT ["sqlermb"]
WORKDIR /root/

