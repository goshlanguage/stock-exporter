FROM cgr.dev/chainguard/go:latest as build

WORKDIR /go/src/github.com/goshlanguage/stock-exporter
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /bin/stock-exporter main.go

RUN go build -o /bin/stock-exporter main.go

ENTRYPOINT ["/bin/stock-exporter"]
