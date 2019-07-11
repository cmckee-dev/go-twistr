package main

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

const (
	errVariableNotSet = "Environment variable required but not set."

	VERSION = "v0.1.0"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func getEnvConfig() map[string]string {
	envConfig := make(map[string]string)

	twitterConsumerKey, ok := os.LookupEnv("TWITTER_CONSUMER_KEY")
	if !ok {
		log.Fatal(errVariableNotSet)
	}
	envConfig["twitterConsumerKey"] = twitterConsumerKey

	twitterConsumerSecret, ok := os.LookupEnv("TWITTER_CONSUMER_SECRET")
	if !ok {
		log.Fatal(errVariableNotSet)
	}
	envConfig["twitterConsumerSecret"] = twitterConsumerSecret

	twitterAccessToken, ok := os.LookupEnv("TWITTER_ACCESS_TOKEN") if !ok {
		log.Fatal(errVariableNotSet)
	}
	envConfig["twitterAccessToken"] = twitterAccessToken

	twitterAccessTokenSecret, ok := os.LookupEnv("TWITTER_ACCESS_TOKEN_SECRET")
	if !ok {
		log.Fatal(errVariableNotSet)
	}
	envConfig["twitterAccessTokenSecret"] = twitterAccessTokenSecret

	return envConfig
}

func getTwitterClient(creds *Credentials) (*twitter.Client, error) {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)

	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("Users Account:\n%+v\n", user)
	return client, nil
}

func main() {
	log.Println("Starting Twistr")
	log.Println("The Go-Twitter-Stream")
	log.Println(VERSION)

	envConfig := getEnvConfig

	log.Println(envConfig)

	// creds := Credentials{
	// 	AccessToken:       envConfig["twitterAccessToken"],
	// 	AccessTokenSecret: envConfig["twitterAccessTokenSecret"],
	// 	ConsumerKey:       envConfig["twitterConsumerKey"],
	// 	ConsumerSecret:    envConfig["twitterConsumerSecret"],
	// }

	// log.Printf("%+v\n", creds)
}
