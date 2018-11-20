FROM golang:1.10-alpine

RUN apk add --no-cache git \
    curl

WORKDIR /go/src/github.com/hiko1129/qiicon
COPY . .

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
ENV PATH $PATH:/go/bin
RUN dep ensure