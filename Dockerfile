FROM golang

ADD . /go/src/github.com/saromanov/user-srv
RUN ./build.sh
ENTRYPOINT /go/bin/user-srv
