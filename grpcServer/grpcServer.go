package grpcServer

import (
	"fmt"
	"github.com/apix76/ShortenURL/Conf"
	"github.com/apix76/ShortenURL/Db/DbAccess"
	"github.com/apix76/ShortenURL/Proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func GrpcServer(conf Conf.Conf) {

	db, err := DbAccess.Db.NewDb(conf.PgsqlNameServe)
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp", conf.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}

	serv := Proto.Server{Db: db}
	grpcServer := grpc.NewServer()
	Proto.RegisterShortenURLServer(grpcServer, &serv)

	fmt.Println("Start grpcServer server.")
	log.Fatal(grpcServer.Serve(l))
}
