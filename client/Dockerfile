# syntax=docker/dockerfile:1

# build stage
FROM node:14.20-alpine AS build

WORKDIR /app

COPY package*.json /app/

RUN npm install 

COPY . .

ENV VUE_APP_API_URL=http://localhost:3000

RUN npm run build


# prod stage
FROM nginx:stable-alpine AS prod

COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80 

CMD ["nginx", "-g", "daemon off;"]