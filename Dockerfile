FROM golang:1.17 as build_base
RUN mkdir /app
WORKDIR /app

# Force the go compiler to use modules
ENV GO111MODULE=on
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# This image builds the weavaite server
FROM build_base AS server_builder
# Here we copy the rest of the source code
COPY *.go ./
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o /demo -tags netgo -ldflags '-w -extldflags "-static"' .

### Put the binary onto base image
FROM heroku/heroku:16
LABEL maintainer="DarkCorNer <darkcorner17375@gmail.com>"
EXPOSE 8080
COPY --from=server_builder /demo /demo
CMD ["/demo"]