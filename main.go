package main

import (
	"github.com/apix76/ShortenURL/Conf"
	"github.com/apix76/ShortenURL/grpcServer"
	"github.com/apix76/ShortenURL/rest"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf, err := Conf.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	if conf.GrpcPort == "" && conf.GrpcPort == "" {
		log.Print("You didn't specify ports")
		return
	}

	if conf.HttpPort != "" {
		go rest.Http(conf)
	}

	if conf.GrpcPort != "" {
		grpcServer.GrpcServer(conf)
	}

}
