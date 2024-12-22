# Build container
FROM node:22-alpine3.20 as build
RUN apk update && apk upgrade
WORKDIR /app
COPY . /app
RUN npm install && npm run build

# Run container
FROM nginx
WORKDIR /app

COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/build /usr/share/nginx/html

EXPOSE 8080
ENV HOST=0.0.0.0 PORT=8080 NODE_ENV=production
