package rest

import (
	"encoding/json"
	"fmt"
	"github.com/apix76/ShortenURL/Conf"
	"github.com/apix76/ShortenURL/Db"
	"github.com/apix76/ShortenURL/Db/DbInterface"
	"github.com/apix76/ShortenURL/UseCase"
	"log"
	"net"
	"net/http"
)

type HTTPHandler struct{ Db DbInterface.Db }

func Http(conf Conf.Conf) {
	mux := Handler(conf.PgsqlNameServe)
	l, err := net.Listen("tcp", conf.HttpPort)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.Serve(l, mux))

}

func Handler(PgsqlNameServe string) http.Handler {

	db, err := Db.NewDb(PgsqlNameServe)
	if err != nil {
		log.Fatal(err)
	}

	handler := HTTPHandler{Db: db}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler.ServeGet)   //TODO: принимает только методы get
	mux.HandleFunc("POST /", handler.ServePost) //TODO: принимает только методы post
	fmt.Println("Start http server")

	return mux
}

func (p *HTTPHandler) ServePost(w http.ResponseWriter, req *http.Request) {
	type Url struct {
		URL string `json:"URL"`
	}

	type ShortenUrl struct {
		ShortURL string `json:"ShortURL"`
	}

	UrlFromRequest := Url{}
	Resp := ShortenUrl{}

	err := json.NewDecoder(req.Body).Decode(&UrlFromRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	defer req.Body.Close()

	if UrlFromRequest.URL == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("url is empty"))
	}

	Resp.ShortURL = UseCase.ShortenURL(UrlFromRequest.URL)

	if _, err := p.Db.Get(Resp.ShortURL); err != nil {
		if err == DbInterface.ErrNoExist {
			if err = p.Db.Add(Resp.ShortURL, UrlFromRequest.URL); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(Resp)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *HTTPHandler) ServeGet(w http.ResponseWriter, req *http.Request) {
	type ShortenUrl struct {
		ShortURL string `json:"ShortURL"`
	}

	type Url struct {
		URL string `json:"URL"`
	}

	RequestURL := ShortenUrl{}
	ResponseURL := Url{}

	err := json.NewDecoder(req.Body).Decode(&RequestURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	defer req.Body.Close()

	if RequestURL.ShortURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`ShortURL is empty`))
	}

	if ResponseURL.URL, err = g.Db.Get(RequestURL.ShortURL); err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`Url Not Found`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(ResponseURL)
	if err != nil {
		log.Fatal(err)
	}
}
