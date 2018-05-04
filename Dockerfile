FROM golang:onbuild
EXPOSE 8080
ENV GOBIN /go/bin

# build directories
RUN mkdir /app
RUN mkdir /go/src/imageService
ADD . /go/src/imageService
WORKDIR /go/src/imageService

# Go dep!
RUN go get -u github.com/golang/dep/...
RUN dep ensure
#
## Build my app
RUN go build -o /imageService/main .
CMD ["/imageService/main"]