FROM golang:1.11.2



RUN go get -u github.com/golang/dep/cmd/dep

RUN mkdir -p /go/src/github.com/JYama321/gin-crud-template

WORKDIR /go/src/github.com/JYama321/gin-crud-template

COPY Gopkg.lock  .
COPY Gopkg.toml .

RUN dep ensure -v -vendor-only=true

COPY . /go/src/github.com/JYama321/gin-crud-template

RUN ls /go/src/github.com/JYama321/gin-crud-template/vendor/github.com/jinzhu/gorm
RUN dep ensure

EXPOSE 3000