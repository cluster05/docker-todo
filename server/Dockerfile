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

ENV ENV=$ENV 
ENV PORT=$PORT 
ENV ClientAppURL=$ClientAppURL 
ENV MongoDNS=$MongoDNS
ENV MongoDB=$MongoDB

EXPOSE 3000

CMD [ "./server" ]
