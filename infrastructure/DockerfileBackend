# Build container
FROM golang:1.23.4-alpine AS build
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o ./godocker
 
# Deploy container
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/godocker .
RUN apk --no-cache add sqlite

EXPOSE 1337
CMD [ "./godocker" ]
