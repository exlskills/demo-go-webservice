# Part 1 (Builder)
FROM golang:1.11 as gobuilder

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# TODO Replace `exlskills` with your GitHub username!
COPY . /go/src/github.com/exlskills/demo-go-webservice
WORKDIR /go/src/github.com/exlskills/demo-go-webservice

# Install dependencies
RUN dep ensure -v

RUN go build

# Part 2 (Final)
FROM alpine:3.7

COPY --from=gobuilder /go/src/github.com/exlskills/demo-go-webservice/demo-go-webservice /home/demo-go-webservice

ENTRYPOINT /home/demo-go-webservice
EXPOSE 3333

