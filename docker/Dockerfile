FROM golang:1.14.2-stretch AS builder

RUN apt-get update && apt-get install -y git

# env for go workdir and mod vendor
ENV GO_WORKDIR /workdir
ENV GOFLAGS=-mod=vendor

# claim workdir and move to workdir loc
WORKDIR $GO_WORKDIR

# copy files into workdir
ADD . $GO_WORKDIR
# RUN go get # use vendor folder instead of update each time

# 使用 -X main.gitcommitnum 動態將 `git_commit_num` 帶入二進制參數
RUN go build -ldflags "-X main.version=`git rev-parse --short=6 HEAD`" -o servicebin

FROM centos:7
ENV GO_WORKDIR /workdir

# 將 /app 加入 container 的環境路徑，啟動後可以傳入指令
ENV PATH /app
WORKDIR /app

# copy binary into container (從 builder 的目錄下將執行檔複製到 centos 內)
COPY --from=builder $GO_WORKDIR/servicebin servicebin

# COPY ./conf.d ./conf.d
# COPY env.yaml .

# CMD ["./servicebin","--config","./env.yaml"]