package main

import (
	"fmt"
	"github.com/bsdf/twitter"
	"io/ioutil"
	"os"
)

func main() {
	t := twitter.Twitter{
		ConsumerKey:      "",
		ConsumerSecret:   "",
		OAuthToken:       "",
		OAuthTokenSecret: "",
	}

	bytes, _ := ioutil.ReadAll(os.Stdin)
	_, err := t.Tweet(string(bytes))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
