package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
)

func Round(f float64) float64 {
	return math.Floor(f + .5)
}

func RoundPlus(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return Round(f*shift) / shift
}

func CacheFaves(twapi *anaconda.TwitterApi) {

	faveStringValues := url.Values{}
	faveStringValues.Set("count", "20")
	faveGetResults, _ := twapi.GetFavorites(faveStringValues)
	for _, tweet := range faveGetResults {
		fmt.Println(tweet.Text, tweet.Id)
	}
}

func GetFaveCount(twapi *anaconda.TwitterApi, sname string) int {
	userQueryValues := url.Values{}
	userQueryValues.Set("screen_name", sname)
	userq, e := twapi.GetUsersShow(sname, userQueryValues)
	if e != nil {
		fmt.Printf("Twitter User Query Error: %v\n", e)

	}
	return userq.FavouritesCount
}

func main() {
	// file contains twitter api keys
	const keyFile = "config.json"
	// API max call limit on 15 min reset
	const apiCallLimit = 15
	// 200 count max per call
	const apiFaveCallCount = 200
	const screenName = "MarkPashby"
	apiKeys := LoadAuthKeys(keyFile)
	twApi := PerformAuth(apiKeys)
	userFaveCount := GetFaveCount(twApi, screenName)
	CacheFaves(twApi)

}
