FROM golang:1.8

RUN mkdir /data
COPY . "$GOPATH/src/github.com/MBControlGroup/MBCG-BE-Login/"
RUN cd "$GOPATH/src/github.com/MBControlGroup/MBCG-BE-Login" && go get -v && go install -v
RUN cd "$GOPATH/src/github.com/MBControlGroup/MBCG-BE-Login/service" && go get -v && go install -v
RUN cd "$GOPATH/src/github.com/MBControlGroup/MBCG-BE-Login/entities" && go get -v && go install -v
RUN cd "$GOPATH/src/github.com/MBControlGroup/MBCG-BE-Login/token" && go get -v && go install -v

EXPOSE 9200

VOLUME /data
