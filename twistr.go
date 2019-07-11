package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

const (
	errVariableNotSet = "environment variable required but not set."

	twitterConsumerKey       = "TWITTER_CONSUMER_KEY"
	twitterConsumerSecret    = "TWITTER_CONSUMER_SECRET"
	twitterAccessToken       = "TWITTER_ACCESS_TOKEN"
	twitterAccessTokenSecret = "TWITTER_ACCESS_TOKEN_SECRET"

	VERSION = "v0.1.0"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func getTwitterCredentials() *Credentials {

	_, ok := os.LookupEnv(twitterConsumerKey)
	if !ok {
		log.Fatalf("%s %s\n", twitterConsumerKey, errVariableNotSet)
	}

	_, ok = os.LookupEnv(twitterConsumerSecret)
	if !ok {
		log.Fatalf("%s %s\n", twitterConsumerSecret, errVariableNotSet)
	}

	_, ok = os.LookupEnv(twitterAccessToken)
	if !ok {
		log.Fatalf("%s %s\n", twitterAccessToken, errVariableNotSet)
	}

	_, ok = os.LookupEnv(twitterAccessTokenSecret)
	if !ok {
		log.Fatalf("%s %s\n", twitterAccessTokenSecret, errVariableNotSet)
	}

	return &Credentials{
		ConsumerKey:       os.Getenv(twitterConsumerKey),
		ConsumerSecret:    os.Getenv(twitterConsumerSecret),
		AccessToken:       os.Getenv(twitterAccessToken),
		AccessTokenSecret: os.Getenv(twitterAccessTokenSecret),
	}
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

func sendTweet(message string, client *twitter.Client) {
	tweet, resp, err := client.Statuses.Update(message, nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)
}

func main() {
	log.Println("Starting Twistr")
	log.Println("The Go-Twitter-Stream")
	log.Println(VERSION)

	creds := Credentials{
		ConsumerKey:       os.Getenv(twitterConsumerKey),
		ConsumerSecret:    os.Getenv(twitterConsumerSecret),
		AccessToken:       os.Getenv(twitterAccessToken),
		AccessTokenSecret: os.Getenv(twitterAccessTokenSecret),
	}

	client, err := getTwitterClient(&creds)

	if err != nil {
		log.Println("Error getting twitter client.")
		log.Println(err)
	}

	fmt.Printf("%+v\n", client)

	// msg := "This is a test tweet from Twistr, a golang project."
	// sendTweet(msg, client)
}
