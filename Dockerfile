FROM golang:1.8

RUN mkdir /data
COPY . "$GOPATH/src/github.com/MBControlGroup/login/"
RUN cd "$GOPATH/src/github.com/MBControlGroup/login" && go get -v && go install -v
RUN cd "$GOPATH/src/github.com/MBControlGroup/login/service" && go get -v && go install -v
RUN cd "$GOPATH/src/github.com/MBControlGroup/login/entities" && go get -v && go install -v
RUN cd "$GOPATH/src/github.com/MBControlGroup/login/token" && go get -v && go install -v

EXPOSE 8080

VOLUME /data
