FROM ubuntu:14.04
MAINTAINER zalemwoo "zalemwoo@gmail.com"

RUN apt-get install git -y

ENV GOROOT /goroot
ENV GOPATH /gopath

ENV PATH $GOROOT/bin:$PATH

RUN mkdir -p /go
ADD out/goget-forward /go

RUN echo "alias goget=\"go get -v -u --insecure\"" > /etc/profile.d/goget.sh

ENTRYPOINT [ "/go/goget-forward" ]

VOLUME /goroot /gopath

EXPORT 80
