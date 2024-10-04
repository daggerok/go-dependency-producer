# go-dependency-producer [![ci](https://github.com/daggerok/go-dependency-producer/actions/workflows/ci.yaml/badge.svg)](https://github.com/daggerok/go-dependency-producer/actions/workflows/ci.yaml)
This repository demonstrates how to develop a library in go, use unit testing and run main as go application...

```bash
brwe reinstall go
#echo 'export GOPATH=$HOME/go' >> ~/.zsh
#echo 'export PATH=$PATH:/path/to/go/bin:$GOPATH/bin' >> ~/.zsh
go env GOPATH
go version

mkdir go-dependency-producer && cd $_
go mod init github.com/daggerok/go-dependency-producer

mkdir lib
touch lib/greet.go # ...and implement
touch lib/greet_test.go # ...and implement
touch main.go # ...and implement

go run main.go
go run main.go Maksimko

go build          main.go && ./main
go build -o greet main.go && ./greet

GOPATH="$(pwd):/home/runner/go/bin" go install .
./bin/go-dependency-producer Max
./bin/go-dependency-producer

mkdir .github/workflows
touch .github/workflows/ci.yaml # ...and implement

# git commit...
# git push...
```
