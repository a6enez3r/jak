# Base image
FROM golang:alpine
# Init Go env
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
# Current dir
WORKDIR /build
# Copy files to work dir
COPY . /build
# Init alpine env
RUN apk add --update make
# Env
RUN make init-env
# Deps
RUN make deps-dev
# Vetting
RUN make vet
# Formatting
RUN make format
# Linting
RUN make lint
# Testing
RUN make test