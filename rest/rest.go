package rest

import (
	"ReductionAPI/Conf"
	"ReductionAPI/Db"
	"ReductionAPI/Db/DbInterface"
	"ReductionAPI/UseCase"
	"database/sql"
	"encoding/json"
	"fmt"
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
		return
	}

	defer req.Body.Close()

	Resp.ShortURL = UseCase.ShortenURL(UrlFromRequest.URL)

	if _, err := p.Db.Get(Resp.ShortURL); err != nil {
		if err == DbInterface.ErrNoExist || err == sql.ErrNoRows {
			if err = p.Db.Add(Resp.ShortURL, UrlFromRequest.URL); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

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
		return
	}

	defer req.Body.Close()

	if ResponseURL.URL, err = g.Db.Get(RequestURL.ShortURL); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(ResponseURL)
	if err != nil {
		log.Fatal(err)
	}
}
