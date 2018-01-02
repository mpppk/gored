FROM golang:1.9.2 AS builder
RUN go get -v github.com/golang/dep/cmd/dep
RUN go get -v github.com/Songmu/goxz/cmd/goxz
RUN go get -v github.com/motemen/gobump/cmd/gobump
RUN go get -v github.com/tcnksm/ghr
RUN go get -v github.com/jteeuwen/go-bindata/...
COPY . /go/src/github.com/mpppk/gored
WORKDIR /go/src/github.com/mpppk/gored
RUN make install -f Makefile

FROM circleci/golang:1.9.2
COPY --from=builder /go/bin/* /go/bin/
RUN git config --global user.email "gored-bot@dammy.com"
RUN git config --global user.name  "gored-bot"
USER root
RUN mkdir /init
RUN chown circleci /init
USER circleci
CMD cd /init && gored init