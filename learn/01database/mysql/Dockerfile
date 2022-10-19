FROM alpine:3.12

WORKDIR $GOPATH/src/app
ADD cmd $GOPATH/src/app/cmd
ADD configset $GOPATH/src/app/configset

CMD cmd/cmd