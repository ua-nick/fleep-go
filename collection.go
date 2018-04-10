package fleep

import (
	"encoding/json"
	"io/ioutil"
)

const CollectionFile = "collection.json"

type Collection struct {
	Type      string   `json:"type"`
	Extension string   `json:"extension"`
	Mime      string   `json:"mime"`
	Offset    int      `json:"offset"`
	Signature []string `json:"signature"`
}

func ParseCollection() []Collection {
	var collection []Collection
	data, _ := ioutil.ReadFile(CollectionFile)
	json.Unmarshal(data, &collection)
	return collection
}
