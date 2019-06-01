package main

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	AccessToken string `json:"access_token"`
}

func loadConfig(filepath string) config {
	bytes, err := ioutil.ReadFile(filepath)
	mustNot(err)

	cfg := config{}
	must(json.Unmarshal(bytes, &cfg))

	return cfg
}
