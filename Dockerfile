FROM golang:1.19-alpine3.16 as builder

# tzdata ensures go can bundle the timezones into time.Time which we use to parse times from the api
RUN apk add ca-certificates tzdata && update-ca-certificates

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY main.go .
COPY proxy proxy
COPY docs docs
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w -extldflags "-static"' -o /proxy main.go


FROM scratch

# copy root certificates so the app can make https calls
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /proxy /proxy

EXPOSE 8020

ENV GIN_MODE=release
CMD ["/proxy"]
