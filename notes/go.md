import private repository
>https://blog.wu-boy.com/2020/03/read-private-module-in-golang/  
>go env -w GOPRIVATE=gitlab.geax.io/demeter  
>git config --global url."https://wayne_shen:PASSWORD@gitlab.geax.io".insteadOf "https://gitlab.geax.io"
>git config --global url."git@github.com:".insteadOf "https://github.com/"
>git config --global url."https://$GITHUB_TOKEN:x-oauth-basic@github.com/".insteadOf "https://github.com/"

install protoc
>sudo apt install -y protobuf-compiler

install protoc-gen-gogofaster
>go get github.com/gogo/protobuf/protoc-gen-gogofaster


protoc --gogofaster_out=. ./master_to_room.proto
protoc --gogofaster_out=plugins=grpc:. ./service.proto

# 1. https://github.com/protocolbuffers/protobuf/releases/download/v3.19.1/protoc-3.19.1-linux-x86_64.zip
# 2. install protoc-3.19.1 to /usr/local/bin or ~/go/bin/protoc
# 2.1 install protoc from apt-get: sudo apt install protobuf-compiler
# 3. install golang source: go get github.com/gogo/protobuf/{protoc-gen-gogofaster,protoc-gen-gogofast,protoc-gen-gofast}