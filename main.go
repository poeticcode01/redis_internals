package main

import (
	"flag"
	"log"

	"github.com/poeticcode01/redis_internals/config"
	"github.com/poeticcode01/redis_internals/server"
)

func setupFlags() {

	flag.StringVar(&config.AppConfig.Host, "host", "0.0.0.0", "host for the redis server")

	flag.IntVar(&config.AppConfig.Port, "port", 7379, "port for the redis server")

	flag.Parse()

}

func main() {
	setupFlags()
	log.Println("Spawning TCP echo server on port :", config.AppConfig.Port)
	server.RunSyncTCPServer()
}
