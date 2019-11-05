FROM golang

RUN mkdir -p /go/src/exe
WORKDIR /go/src/exe

COPY . .

RUN go get


RUN go build -ldflags "-linkmode external -extldflags -static" -o server1 main.go

FROM scratch

ADD ca-certificates.crt /etc/ssl/certs/

COPY --from=0 /go/src/exe/server1 /server2
COPY --from=0 /go/src/exe/config.json /config.json
CMD ["/server2"]