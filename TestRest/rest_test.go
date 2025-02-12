package TestRest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apix76/ShortenURL/Conf"
	"github.com/apix76/ShortenURL/rest"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"
)

type TestStruct struct {
	httpPort      string
	Url           Url
	ShortenURL    ShortenURL
	UrlFromServer Url
}

type Url struct {
	URL string `json:"URL"`
}

type ShortenURL struct {
	ShortURL string `json:"ShortURL"`
}

const (
	domen              = "http://127.0.0.1"
	lengthRandomString = 20
)

func TestRest(t *testing.T) {
	conf, err := NewTestConfig()
	if err != nil {
		t.Error(err)
	}

	go rest.Http(conf)
	time.Sleep(1 * time.Second)

	wg := errgroup.Group{}
	for i := 0; i < 10; i++ {
		wg.Go(RequestLoop)
	}
	if err := wg.Wait(); err != nil {
		t.Errorf("Test failed: %v", err)
	}
}

func RequestLoop() error {
	conf, err := NewTestConfig()
	if err != nil {
		return err
	}

	t := TestStruct{httpPort: conf.HttpPort}
	for i := 0; i < 10; i++ {
		t.Url.URL = RandomString()
		if err := t.Post(); err != nil {
			return err
		}
		if err := t.Get(); err != nil {
			return err
		}
		if t.Url.URL != t.UrlFromServer.URL {
			return errors.New("Mismatch Url")
		}
	}
	return nil
}

func (t *TestStruct) Post() error {
	databyte, err := json.Marshal(t.Url)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", domen+t.httpPort, bytes.NewReader(databyte))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&t.ShortenURL)
	if err != nil {
		return err
	}
	return err
}

func (t *TestStruct) Get() error {
	databyte, err := json.Marshal(t.ShortenURL)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("GET", domen+t.httpPort, bytes.NewReader(databyte))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&t.UrlFromServer)
	if err != nil {
		return err
	}
	return err
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

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, lengthRandomString)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
