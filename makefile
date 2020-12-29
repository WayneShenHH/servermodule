
# require .apib file and define your apis, you should install aglio at first
gen_apib:
	aglio -i ./example.apib -o ./example.html

buildflag:
	go build -ldflags "-X GameService/env.Gitcommitnum=`git rev-parse --short=6 HEAD`"

syncmap:
	syncmap -o "module/nats/flowmap.go" -name flowMap -pkg nats "map[string]chan *flow.Flow"

tool:
	go get -u github.com/a8m/syncmap

build:
	go mod tidy
	go mod vendor
	go build

lint:
	golangci-lint run --fast