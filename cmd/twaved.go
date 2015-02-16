package main

import (
	"fmt"
	"net/url"
)

func main() {
	// file contains twitter api keys
	const keyFile = "config.json"
	// Search String Additional Values
	faveStringValues := url.Values{}
	faveStringValues.Set("count", "20")
	apiKeys := LoadAuthKeys(keyFile)
	twApi := PerformAuth(apiKeys)
	faveGetResults, _ := twApi.GetFavorites(faveStringValues)
	for _, tweet := range faveGetResults {
		fmt.Println(tweet.Text, tweet.Id)
	}
}
