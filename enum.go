package main

import "github.com/yrzs/protoc-gen-enum/command"

func main() {
	command.Write(command.Generate(command.Read()))
}
