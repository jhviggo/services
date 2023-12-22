FROM golang:1.21-alpine3.18
 
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /godocker
 
EXPOSE 8080
 
CMD [ “/godocker” ]
