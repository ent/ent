# Multi stage application builder
FROM golang:1.12.9-alpine

# Install tools required to build the project
RUN apk add --no-cache git gcc musl-dev

# Use public go modules proxy
ENV GOPROXY https://proxy.golang.org

# Install goimports for entc.
RUN go get -u golang.org/x/tools/cmd/goimports

# Copy mod files and install dependencies.
WORKDIR /ent
COPY go.mod go.sum ./
RUN go mod download

# Copy project.
COPY . ./
