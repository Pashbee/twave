package main

import (
	"encoding/json"
	"fmt"
	//	"github.com/ChimeraCoder/anaconda"
	"io/ioutil"
	"os"
)

type Tauth struct {
	consumerKey, consumerSecret, accessToken, accessSecret string
}

type jsonobject struct {
	Name  string
	Value string
}

func LoadAuthKeys(filename string) []string {

	var jsontype []jsonobject
	keyValues := make([]string, 4)
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	json.Unmarshal(file, &jsontype)
	for i := range jsontype {
		keyValues[i] = jsontype[i].Value
	}
	return keyValues
}

func PerformAuth() {

}

func main() {
	apiKeys := LoadAuthKeys("./config.json")
	fmt.Println(apiKeys[0])
}
