# Copyright 2019-present Facebook Inc. All rights reserved.
# This source code is licensed under the Apache 2.0 license found
# in the LICENSE file in the root directory of this source tree.

# Application base
FROM golang:1.13

# Install tools required to build the project
RUN apk add --no-cache git gcc musl-dev

# Copy mod files and install dependencies.
WORKDIR /ent
COPY go.mod go.sum ./
RUN go mod download

# Copy project.
COPY . ./
