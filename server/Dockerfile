# syntax=docker/dockerfile:1

# build stage
FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

COPY . .

RUN go build -o server


# prod stage
FROM alpine:latest as prod 

WORKDIR /app

COPY --from=build /app/server /app/

ENV ENV=prod PORT=3000 ClientAppURL=http://localhost:8080 
ENV MongoDNS=mongodb://mongodb:27017 MongoDB=todo

EXPOSE 3000

CMD [ "./server" ]
