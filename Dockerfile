FROM node:20-alpine as web
WORKDIR /app

COPY frontend/package.json /app/package.json
COPY frontend/yarn.lock /app/yarn.lock
RUN yarn install

COPY frontend /app
RUN yarn build

FROM golang:latest as builder
WORKDIR /app

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# src code
COPY . .
COPY --from=web /app/dist /app/backend/web

# env
ENV GOPROXY=https://goproxy.cn,direct
ENV GOOS=linux
ENV CGO_ENABLED=0

# build
RUN go build -tags netgo -o vortexnotes

FROM alpine:latest
COPY --from=builder /app/vortexnotes /app/vortexnotes
EXPOSE 10060
CMD ["/app/vortexnotes"]
