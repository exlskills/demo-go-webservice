# Part 1 (Builder)
FROM golang:1.11-alpine3.7 as gobuilder

RUN apk add --no-cache curl git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# TODO Replace `exlskills` with your GitHub username!
COPY . /go/src/github.com/exlskills/demo-go-webservice
WORKDIR /go/src/github.com/exlskills/demo-go-webservice

# Install dependencies
RUN dep ensure -v

# Build binary
RUN go build

# Part 2 (Final)
FROM alpine:3.7

# Copy over our binary from the builder step
COPY --from=gobuilder /go/src/github.com/exlskills/demo-go-webservice/demo-go-webservice /home/demo-go-webservice

# Configure entrypoint and expose our service's port (3333)
ENTRYPOINT /home/demo-go-webservice
EXPOSE 3333

