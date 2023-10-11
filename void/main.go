package main

import (
	"void/cmd"
	"void/server"
)

func main() {
	cmd.Execute()
	server.ServerStartup()
}
