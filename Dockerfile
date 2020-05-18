#
# https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
#
# BUILD: docker build -t nsip/admindocker .
# TEST: docker run -it -p8097:8097 -v /var/run/docker.sock:/var/run/docker.sock nsip/admindocker
# RUN: docker run -d -p8097:8097 -v /var/run/docker.sock:/var/run/docker.sock nsip/admindocker
############################
# STEP 1 build executable binary
############################
FROM golang:1.13-stretch as builder
COPY . .
RUN go get
RUN go get github.com/labstack/echo/middleware
RUN go get github.com/nsip/admin-docker
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/admindocker cmd/web/main.go
############################
# STEP 2 build a small image
############################
FROM debian:stretch
COPY --from=builder /go/bin/admindocker /go/bin/admindocker
COPY dashboard/index.html /go/bin/dashboard/index.html
WORKDIR /go/bin
CMD ["/go/bin/admindocker"]
