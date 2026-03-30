buildflag:
	go build -ldflags "-X GameService/env.Gitcommitnum=`git rev-parse --short=6 HEAD`"

syncmap:
	syncmap -o "module/nats/flowmap.go" -name flowMap -pkg nats "map[string]chan *flow.Flow"

tool:
	go get -u github.com/a8m/syncmap

build:
	go mod tidy
	go mod vendor
	go build -o wayneutil cmd/wayneutil/main.go
	cp ./wayneutil ~/go/bin/wayneutil

lint:
	golangci-lint run --fast

install:
	go install ./cmd/wayneutil

zshrc:
	echo "source ${PWD}/.zshrc" > ~/.zshrc
	echo "source ${PWD}/.zshrc_golang" >> ~/.zshrc
	echo "source ${PWD}/.zshrc_git" >> ~/.zshrc
	echo "source ${PWD}/.zshrc_fgw" >> ~/.zshrc
	echo "source ${PWD}/.zshrc_emb" >> ~/.zshrc
	zsh ~/.zshrc