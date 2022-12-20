# demo-echo-v4

```
GOPATH="/Users/hexaforce/go"
go mod init demo-echo-v4
go mod tidy
go build -o app
```

```
docker build . -t demo/echo-v4
docker run -it -p 1323:1323 demo/echo-v4
```

go install github.com/swaggo/swag/cmd/swag@latest

export GOROOT=$HOME/go

export PATH=$PATH:$GOROOT/bin

/Users/hexaforce/go/bin/swag init
