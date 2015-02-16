package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"time"
)

func CacheFaves(twapi *anaconda.TwitterApi) {

	faveStringValues := url.Values{}
	faveStringValues.Set("count", "20")
	faveGetResults, _ := twapi.GetFavorites(faveStringValues)
	for _, tweet := range faveGetResults {
		fmt.Println(tweet.Text, tweet.Id)
	}
}

func main() {
	// file contains twitter api keys
	const keyFile = "config.json"
	// API max call limit on 15 min reset
	const apiCallLimit = 15
	apiKeys := LoadAuthKeys(keyFile)
	twApi := PerformAuth(apiKeys)
	CacheFaves(twApi)
}
