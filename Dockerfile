FROM --platform=linux/amd64 golang:1.17-alpine3.13 AS builder
WORKDIR /src  
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app _cmd/main.go
EXPOSE 6060

FROM --platform=linux/amd64 alpine:3.13  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /src/app .
RUN ls
CMD ["./app"]
