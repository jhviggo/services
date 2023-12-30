# Build container
FROM node:18-alpine3.17 as build
RUN apk update && apk upgrade && adduser -D svelteuser
USER svelteuser
WORKDIR /app
COPY --chown=svelteuser:svelteuser . /app
RUN npm install && npm run build

# Run container
FROM node:18-alpine3.17
RUN apk update && apk upgrade && adduser -D svelteuser
USER svelteuser
WORKDIR /app
COPY --chown=svelteuser:svelteuser --from=build /app/build /app/package.json ./
EXPOSE 8080
ENV HOST=0.0.0.0 PORT=8080 NODE_ENV=production
CMD ["node","index.js"]
