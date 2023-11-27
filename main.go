package main

import (
	"termvoid/cmd"
	"termvoid/server"
)

func main() {
	server.ServerStartup()
	cmd.Execute()
}
