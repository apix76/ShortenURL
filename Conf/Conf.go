package Conf

import (
	"encoding/json"
	"os"
)

type Conf struct {
	HttpPort       string
	GrpcPort       string
	PgsqlNameServe string
}

func NewConfig() (Conf, error) {
	var conf Conf

	file, err := os.Open("config.cfg")
	if err != nil {
		return Conf{}, err
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		return Conf{}, err
	}

	return conf, err
}
