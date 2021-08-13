package main

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)
	atlassian.Auth.SetUserAgent("curl/7.54.0")

	var (
		boardID   = 4
		states    = []string{"future", "active"} // valid values: future, active, closed
		startAt   = 0
		maxResult = 50
	)

	sprints, response, err := atlassian.Agile.Board.Sprints(context.Background(), boardID, startAt, maxResult, states)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", response.Bytes.String())
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.Code)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, sprint := range sprints.Values {
		log.Println(sprint)
	}

	fmt.Println(response.Bytes.String())
}