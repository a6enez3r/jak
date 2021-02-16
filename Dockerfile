# Base image
FROM golang:alpine
# Linux depdendencies
RUN apk update
RUN apk upgrade
# bash
RUN apk add bash
# make
RUN apk add --update make
# gotty
ENV GOTTY_BINARY https://github.com/yudai/gotty/releases/download/v1.0.1/gotty_linux_386.tar.gz
# extract
RUN wget $GOTTY_BINARY -O gotty.tar.gz && \
    tar -xzf gotty.tar.gz -C /usr/local/bin/ && \
    rm gotty.tar.gz && \
    chmod +x /usr/local/bin/gotty
# Set necessary go environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
# Move to working directory /build
WORKDIR /build
# Copy files to working directory /build
COPY . /build
# Copy and download dependency using go mod
COPY go.mod .
RUN go mod download
# Install linter
RUN go get -u golang.org/x/lint/golint
# Run tests
#RUN make test
# Run linting
RUN make vet
RUN make lint
# Export necessary port
EXPOSE 8080
# Command to run when starting the container
CMD ["make", "web"]