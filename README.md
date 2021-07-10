# GoLang Practice
This repository has the Go Lang pratice files arranged in chapters according to "The Go Programming language" book by Brian Kernighan.

**Basic Go Cmds:**

1. Compiling Go programs `$ go build path/code.go`

This creates a `code.exe` file on the directory where this cmd is ran and can be executed by `$ code.exe`

2. Running small go files (mostly while development) `go run path/file.exe`

3. Auto-Formatting go files `gofmt [flags] [path]` More details here: https://golang.org/cmd/gofmt/

4. Formatting all files in a project folder
- [Optional] List all files that differ from go format `gofmt -l .`
- Format recursively all files using `gofmt -w .` OR `gofmt -w foldername`