package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main() {

	/*
		----------- Set an environment variable in git bash -----------
		export HOST="https://ctreminiom.atlassian.net/"
		export MAIL="MAIL_ADDRESS"
		export TOKEN="TOKEN_API"

		Docs: https://stackoverflow.com/questions/34169721/set-an-environment-variable-in-git-bash
	*/

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	payload := &jira.PermissionCheckPayload{
		GlobalPermissions: []string{"ADMINISTER"},
		AccountID:         "", //
		ProjectPermissions: []*jira.BulkProjectPermissionsScheme{
			{
				Issues:      nil,
				Projects:    []int{10000},
				Permissions: []string{"EDIT_ISSUES"},
			},
		},
	}

	grants, response, err := atlassian.Permission.Check(context.Background(), payload)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, permission := range grants.ProjectPermissions {
		log.Println(permission.Permission, permission.Issues)
	}

}