FROM circleci/golang:1.9.2
RUN go get -v github.com/golang/dep/cmd/dep
RUN go get -v github.com/Songmu/goxz/cmd/goxz
RUN go get -v github.com/motemen/gobump/cmd/gobump
RUN go get -v github.com/tcnksm/ghr
RUN git config --global user.email "goreleng-bot@dammy.com"
RUN git config --global user.name  "goreleng-bot"
COPY ./Makefile /Makefile