package main

import "fmt"
import "os"

func main() {
	fmt.Println("hello world")

	args := os.Args[1:]
	fmt.Println(args)
}

// go version > 1.23
// go run 01.go
// go build 01.go

// gofmt 01.go
// go fmt 01.go
// go mod init <name>
// go fmt .

// go install golang.org/x/tools/cmd/goimports@latest