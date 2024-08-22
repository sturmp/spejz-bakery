FROM --platform=linux/arm64 golang:alpine as build

ENV GO111MODULE=on

RUN apk update && apk add --no-cache git build-base

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