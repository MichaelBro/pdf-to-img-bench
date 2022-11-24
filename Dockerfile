FROM centos:7

RUN yum -y update
RUN yum install gcc git libmpfr4 -y && yum clean all
ENV GO_VERSION 1.19.3
RUN curl -L https://golang.org/dl/go$GO_VERSION.linux-amd64.tar.gz -o go$GO_VERSION.linux-amd64.tar.gz && tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz && rm -f go$GO_VERSION.linux-amd64.tar.gz
ENV PATH /usr/local/go/bin:$PATH
ENV GOROOT /usr/local/go

ENV PDFIUM_VERSION 5431
RUN curl -L https://github.com/bblanchon/pdfium-binaries/releases/download/chromium%2F$PDFIUM_VERSION/pdfium-linux-x64.tgz -o pdfium.tar.gz && mkdir "/home/pdfium" && tar -C /home/pdfium -xzf pdfium.tar.gz && rm -f pdfium.tar.gz

COPY pdfium.pc /usr/lib64/pkgconfig/pdfium.pc
RUN cp /home/pdfium/lib/libpdfium.so /usr/lib64/libpdfium.so

RUN yum -y install java-11-openjdk-devel which #for pdfbox

RUN yum -y install graphviz

COPY bench /home/bench
RUN mkdir "/home/files"

WORKDIR /home/bench

RUN go mod tidy
RUN go build

