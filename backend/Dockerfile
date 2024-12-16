# Build container
FROM golang:1.22-alpine3.20 AS build
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o ./godocker
 
# Deploy container
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/godocker .

EXPOSE 8080
CMD [ "./godocker" ]