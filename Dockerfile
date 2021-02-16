FROM golang:alpine as builder
WORKDIR /stocksync
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o stocksync cmd/*.go

FROM scratch
COPY --from=builder /stocksync/stocksync .
ENTRYPOINT ["./stocksync","http-serve"]