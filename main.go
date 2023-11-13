package main

import (
	"termvoid/cmd"
	"termvoid/server"
)

func main() {
	cmd.Execute()
	server.ServerStartup()
}
