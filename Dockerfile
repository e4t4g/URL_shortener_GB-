FROM golang:1.17 AS builder

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .

COPY /bin/url_shortener_gb /bin/url_shortener_gb
COPY cmd/favicon.ico /favicon.ico
COPY web/create/ /web/create/template.html
COPY web/result/ /web/result/result.html
COPY web/stat/ /web/stat/stat.html
COPY cmd/configs/app.yaml configs/app.yaml

RUN go build -o /bin/url_shortener_gb ./cmd/

EXPOSE 8080

CMD ["/bin/url_shortener_gb"]


