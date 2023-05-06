# How To Install?

## MacOs

```sh
// TODO: brew install
```

## From Binary Releases (macOS, Windows, Linux)

## Using [`GO`](https://go.dev/dl) toolchain

```sh
go install github.com/heybran/todo-app@latest
```

## Build The Binary Yourself

# building the program for intel macs
GOOS=darwin GOARCH=amd64 go build -o gurl-mac-amd64 cmd/gurl/main.go 
# building the program for M1 macs
GOOS=darwin GOARCH=arm64 go build -o gurl-mac-arm64 cmd/gurl/main.go 
# building the program for 64 bits amd/intel linux
GOOS=linux GOARCH=amd64 go build -o gurl-linux-amd64 cmd/gurl/main.go