FROM golang:1.11.2



RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/JYama321/penguin_bbs

COPY Gopkg.lock  .
COPY Gopkg.toml .

RUN dep ensure -v -vendor-only=true

COPY . /go/src/github.com/JYama321/penguin_bbs

RUN ls /go/src/github.com/JYama321
RUN dep ensure

EXPOSE 3000