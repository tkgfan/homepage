FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build
COPY . .
RUN go build -ldflags "-s -w" -o /app/app main.go

FROM alpine:latest

WORKDIR /app
ENV TZ Asia/Shanghai
COPY --from=builder /build/web /app/web
COPY --from=builder /app/app /app/app

EXPOSE 8888

CMD ["./app"]
