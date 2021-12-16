import private repository
>https://blog.wu-boy.com/2020/03/read-private-module-in-golang/  
>go env -w GOPRIVATE=gitlab.geax.io/demeter  
>git config --global url."https://wayne_shen:PASSWORD@gitlab.geax.io".insteadOf "https://gitlab.geax.io"

install protoc
>sudo apt install -y protobuf-compiler

install protoc-gen-gogofaster
>go get github.com/gogo/protobuf/protoc-gen-gogofaster