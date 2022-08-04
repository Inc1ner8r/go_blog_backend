FROM golang:alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN apk add build-base
RUN go mod download
COPY . .
RUN apk add --no-cache git && go build -o engine . && apk del git

FROM alpine
WORKDIR /app
COPY --from=builder /app/engine .

EXPOSE 8080

CMD [ "./engine" ]