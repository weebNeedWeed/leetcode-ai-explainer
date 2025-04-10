FROM node:22.14-bookworm AS build

WORKDIR /app

COPY ./web/package*.json ./

RUN npm ci

COPY ./web/ ./

RUN npm run build

###
FROM nginx:1.27.4-alpine

COPY --from=build /app/dist/ /usr/share/nginx/html/

COPY ./web/nginx.conf /etc/nginx/

EXPOSE 80

