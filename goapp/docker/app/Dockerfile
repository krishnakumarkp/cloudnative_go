#Build environment
#------------------
FROM golang:1.13-alpine as build-env
WORKDIR /files/goapp
RUN apk update && apk add --no-cache gcc musl-dev git bash
COPY . .
RUN go mod download
RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate

#Deployment environment
#----------------------
FROM alpine
RUN apk update && apk add --no-cache bash mysql-client

COPY --from=build-env /files/goapp/bin/app /goapp/
COPY --from=build-env /files/goapp/bin/migrate /goapp/
COPY --from=build-env /files/goapp/migrations /goapp/migrations
COPY --from=build-env /files/goapp/docker/app/bin /usr/local/bin/myapp/
RUN chmod +x /usr/local/bin/myapp/*
CMD ["/goapp/app"]
EXPOSE 8080