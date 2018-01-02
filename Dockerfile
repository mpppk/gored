FROM circleci/golang:1.9.2
RUN go get -v github.com/golang/dep/cmd/dep
RUN go get -v github.com/Songmu/goxz/cmd/goxz
RUN go get -v github.com/motemen/gobump/cmd/gobump
RUN go get -v github.com/tcnksm/ghr
RUN go get -v github.com/jteeuwen/go-bindata/...
RUN git config --global user.email "gored-bot@dammy.com"
RUN git config --global user.name  "gored-bot"
COPY --chown=circleci . /go/src/github.com/mpppk/gored
WORKDIR /go/src/github.com/mpppk/gored
RUN make install -f Makefile.gored
USER root
RUN mkdir /init
RUN chown circleci /init
USER circleci
CMD cd /init && gored init