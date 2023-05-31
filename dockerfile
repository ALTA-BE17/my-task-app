FROM golang:alpine AS builder
WORKDIR /app/
COPY . .
RUN go build -o /app/main /app/main.go

FROM alpine:latest 
WORKDIR /app/
COPY --from=builder /app/ /app/
CMD [ "/app/main" ]