package main

import (
	"github.com/apix76/ShortenURL/Conf"
	"github.com/apix76/ShortenURL/Servers/grpcServer"
	"github.com/apix76/ShortenURL/Servers/rest"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf, err := Conf.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	if conf.GrpcPort == "" && conf.HttpPort == "" {
		log.Print("You didn't specify ports")
		return
	}

	if conf.HttpPort != "" && conf.GrpcPort != "" {
		go rest.Http(conf)
	} else if conf.HttpPort != "" && conf.GrpcPort == "" {
		rest.Http(conf)
	}

	if conf.GrpcPort != "" {
		grpcServer.GrpcServer(conf)
	}
}
