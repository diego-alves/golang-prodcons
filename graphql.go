package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type GraphQL struct {
	Query     string            `json:"query"`
	Variables map[string]string `json:"variables"`
}

func sendRequest(url string, query string, variables map[string]string) []byte {
	obj := GraphQL{query, variables}
	objJson, _ := json.Marshal(obj)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(objJson))
	req.Header.Add("Authorization", "token "+os.Getenv("GITHUB_TOKEN"))
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body
}
