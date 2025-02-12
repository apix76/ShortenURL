package TestGrpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apix76/ShortenURL/Conf"
	"github.com/apix76/ShortenURL/Proto"
	"github.com/apix76/ShortenURL/grpcServer"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

type Execute struct {
	client Proto.ShortenURLClient
}

const lengthRandomString = 20

func TestGrpc(t *testing.T) {
	conf, err := NewTestConfig()
	if err != nil {
		t.Error(err)
	}

	go grpcServer.GrpcServer(conf)
	time.Sleep(1 * time.Second)

	con, err := grpc.NewClient(fmt.Sprintf("127.0.0.1%s", conf.GrpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	e := Execute{client: Proto.NewShortenURLClient(con)}

	wg := errgroup.Group{}
	for i := 0; i < 10; i++ {
		wg.Go(e.Testing)
	}

	if err := wg.Wait(); err != nil {
		t.Error(err)
	}
}

func (e *Execute) Testing() error {
	for i := 0; i < 10; i++ {
		url := RandomString()
		shorturl, err := e.client.GetShortenURL(context.Background(), &Proto.URL{Url: url})
		if err != nil {
			return err
		}
		UrlFromServer, err := e.client.GetAllURL(context.Background(), shorturl)
		if err != nil {
			return err
		}

		if UrlFromServer.Url != url {
			errors.New("Mismatch url")
		}
	}
	return nil
}

func NewTestConfig() (Conf.Conf, error) {
	var conf Conf.Conf

	file, err := os.Open("Testconfig.cfg")
	if err != nil {
		return Conf.Conf{}, err
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		return Conf.Conf{}, err
	}

	return conf, err
}

func RandomString() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, lengthRandomString)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
