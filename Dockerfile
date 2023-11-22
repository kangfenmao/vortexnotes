FROM node:20-alpine as web
WORKDIR /app
COPY frontend /app
RUN yarn install && yarn build

FROM golang:latest as builder
WORKDIR /app
COPY . .
COPY --from=web /app/dist /app/backend/web
ENV GOOS=linux
ENV CGO_ENABLED=0
RUN go build -tags netgo -o vortexnotes

FROM alpine:latest
COPY --from=builder /app/vortexnotes /app/vortexnotes
EXPOSE 7701
CMD ["/app/vortexnotes"]
