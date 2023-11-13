FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o vortexnotes

FROM alpine:latest
COPY --from=builder /app/vortexnotes /app/vortexnotes
EXPOSE 6480

CMD ["/app/vortexnotes"]
