package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/0xAX/notificator"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var notify *notificator.Notificator

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Usage: %v <owner> <repo> <issueNumber>\n", os.Args[0])
		os.Exit(0)
	}

	accessToken, exists := os.LookupEnv("GITHUB_ACCESS_TOKEN")

	if exists != true {
		fmt.Println("Warning: no GITHUB_ACCESS_TOKEN environment variable detected, will only be able to check public repos")
	}

	owner := os.Args[1]
	repo := os.Args[2]
	prNumber, err := strconv.Atoi(os.Args[3])

	notify = notificator.New(notificator.Options{
		DefaultIcon: "",
		AppName:     "Can I push?",
	})

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	if err != nil {
		panic(err)
	}

	tick := time.Tick(3 * time.Second)

	for {
		select {
		case <-tick:
			pr, response, err := client.PullRequests.Get(ctx, owner, repo, prNumber)

			if err != nil {
				fmt.Println(fmt.Errorf("%v", err))
				panic(err)
			}

			if response.StatusCode != http.StatusOK {
				panic(response.Status)
			}

			if *pr.State == "closed" {
				fmt.Println("Already merged")
				return
			}

			ref := pr.GetHead().GetRef()

			statuses, _, err := client.Repositories.GetCombinedStatus(ctx, owner, repo, ref, nil)

			if err != nil {
				fmt.Println(fmt.Errorf("%v", err))
				panic(err)
			}

			if *pr.Mergeable == true && *statuses.State == "success" {
				message := fmt.Sprintf("Your PR %s/%s#%d is ready to be merged!", owner, repo, prNumber)
				notify.Push("You can push!", message, "", notificator.UR_NORMAL)
				return
			}
		}
	}
}
