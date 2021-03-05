FROM ubuntu:18.04

ENV LANG C.UTF-8

RUN echo "" > /etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse" >> /etc/apt/sources.list \
    && apt-get clean \
    && apt-get update

RUN DEBIAN_FRONTEND=noninteractive \
    && TZ=Asia/Shanghai \
    && apt-get install -y vim \
    && apt-get install -y wget \
    && apt-get install -y dpkg \
    && apt-get install -y unzip \
    && apt-get autoremove -y

RUN apt-get install -y git

RUN wget -c https://studygolang.com/dl/golang/go1.15.3.linux-amd64.tar.gz
RUN tar -C /usr/local -xvf go1.15.3.linux-amd64.tar.gz

ENV PATH $PATH:/usr/local/go/bin

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

ADD . /src
WORKDIR /src
