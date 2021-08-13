package main

import (
	"context"
	"encoding/json"
	"github.com/ctreminiom/go-atlassian/jira"
	"github.com/ctreminiom/go-atlassian/jira/agile"
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
		options = &agile.IssueOptionScheme{
			JQL:           "project = KP",
			ValidateQuery: true,
			Fields:        []string{"status", "issuetype", "summary"},
			Expand:        []string{"changelog"},
		}

		epicKey   = "KP-16"
		startAt   = 0
		maxResult = 50
	)

	issues, response, err := atlassian.Agile.Epic.Issues(context.Background(), epicKey, startAt, maxResult, options)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", response.Bytes.String())
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.Code)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, issue := range issues.Issues {
		log.Println(issue.Key)
	}

	//If you want to extract the field metadata
	var searchPage jira.IssueSearchScheme
	if err = json.Unmarshal(response.Bytes.Bytes(), &searchPage); err != nil {
		log.Fatal(err)
	}

	for _, issue := range searchPage.Issues {
		log.Println(issue.Key, issue.Fields.Summary, issue.Fields.IssueType.Name)
	}
}