## Create network first

> docker network create todo_network

## Create Images For MongoDB,Server,Client


### Pull mongo image from https://hub.docker.com

> docker pull mongo

### Create new Image for golang application 

> cd server
> docker build -t server .


### Create new Image for vue application

> cd client
> docker build -t client .

## Build Container

### Up Mongo Conainer  with [ named volumes (preserve db data), network ]

> docker run --name mongodb --rm --network todo_network -d -v data:/data/db mongo

### Up Server Conainer with [ network ]

> docker run --name server --rm --network todo_network -d -p 3000:3000 server:latest

### Up Client Conainer with [ network ]

> docker run --name client --rm --network todo_network -d -p 8080:8080 client:latest


