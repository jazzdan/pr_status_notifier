package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	owner := os.Args[1]
	repo := os.Args[2]
	prNumber, err := strconv.ParseInt(os.Args[3], 10, 64)

	if err != nil {
		panic(err)
	}

	tick := time.Tick(5 * time.Second)

	for {
		select {
		case <-tick:
			pr, _, err := client.PullRequests.Get(ctx, owner, repo, int(prNumber))

			if err != nil {
				fmt.Println(fmt.Errorf("%v", err))
				panic(err)
			}

			if *pr.Mergeable == true {
				fmt.Println("Mergeable!")
				return // I expect it to break out of the for loop here
			}
		}
	}
}
