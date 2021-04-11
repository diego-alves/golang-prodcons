package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type GraphQL struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables"`
}

func Request(query string) []byte {
	return RequestE(query, nil)
}

func RequestE(query string, variables map[string]interface{}) []byte {
	obj := GraphQL{query, variables}
	objJson, _ := json.Marshal(obj)
	req, _ := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(objJson))
	req.Header.Add("Authorization", "token "+os.Getenv("GITHUB_TOKEN"))
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body
}
