package main

import (
	"pharos/services/cli"
	"pharos/services/conf"
	"pharos/services/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cli.Parse()
	server.Start(conf.NewConfig("dev"))
}
