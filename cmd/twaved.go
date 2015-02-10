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

func PerformAuth() {
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))
	var jsontype []jsonobject
	json.Unmarshal(file, &jsontype)
	fmt.Printf("%+v", jsontype)
}

func main() {
	PerformAuth()
}
