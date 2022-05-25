FROM golang:1.18 AS builder

COPY . /github.com/e4t4g/URL_shortener_GB-/
WORKDIR /github.com/e4t4g/URL_shortener_GB-/

RUN go mod download
RUN go build -o /bin/url_shortener_gb ./cmd

FROM alpine:latest

COPY --from=0 /github.com/e4t4g/URL_shortener_GB-/bin/url_shortener_gb /bin/url_shortener_gb
COPY --from=0 /github.com/e4t4g/URL_shortener_GB-/cmd/configs/app.yaml configs/app.yaml

CMD ["./bin/url_shortener_gb"]

