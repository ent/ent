# Copyright 2019-present Facebook Inc. All rights reserved.
# This source code is licensed under the Apache 2.0 license found
# in the LICENSE file in the root directory of this source tree.

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
