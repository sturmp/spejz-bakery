FROM golang:alpine AS build

ENV GO111MODULE=on

RUN apk add --no-cache build-base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go install github.com/mattn/go-sqlite3

COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o /app/api

FROM alpine
COPY --from=build /app/api /api
COPY --from=build /app/appconfig.json /appconfig.json
COPY --from=build /app/migration /migration
CMD [ "./api" ]