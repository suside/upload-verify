package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"strings"
)

var cmp comparator

type comparator struct {
	verbose          *bool
	functionSelected cmpFunction
	functionsAvaible map[string]cmpFunction
	httpClient       *http.Client
	filesCheckCount  int
}

type cmpFunction func(os.FileInfo) string

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cmp = comparator{
		filesCheckCount:  0,
		httpClient:       &http.Client{Transport: tr},
		functionsAvaible: make(map[string]cmpFunction),
	}
}

func (c *comparator) register(function cmpFunction) {
	functionPointer := reflect.ValueOf(function).Pointer()
	functionName := strings.Split(runtime.FuncForPC(functionPointer).Name(), ".")[1]
	c.functionsAvaible[functionName] = function
}

func (c *comparator) useCmpFunction(function string) bool {
	if c.functionsAvaible[function] != nil {
		c.functionSelected = c.functionsAvaible[function]
		return true
	}
	return false
}

func (c *comparator) String() string {
	var keys []string
	for key := range c.functionsAvaible {
		keys = append(keys, key)
	}
	return strings.Join(keys, ", ")
}

func (c *comparator) assertURLFileMatchLocalPath(url string, path string) {
	c.filesCheckCount++
	response, err := c.httpClient.Head(url) // TODO perhpas full body sha1?
	if err != nil {
		log.Fatalln(err)
	}
	if response.StatusCode != 200 {
		log.Fatalf("Status %d unexcpected for %s", response.StatusCode, path)
	}
	defer response.Body.Close()
	stat, err := os.Stat(path)
	if err != nil {
		log.Fatalln(err)
	}
	etagHeader := response.Header.Get("Etag") // TODO support mod time?
	if etagHeader != c.functionSelected(stat) {
		log.Fatalf("FATAL! %s remote:%s local:%s", path, etagHeader, c.functionSelected(stat))
	} else if *c.verbose {
		log.Printf("OK %s remote:%s local:%s", path, etagHeader, c.functionSelected(stat))
	}
}
