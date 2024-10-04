# go-dependency-producer

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

# git commit...
# git push...
```
