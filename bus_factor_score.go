package main

import (
	"context"
	"fmt"
	"log"
    "strconv"
	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

func bus_factor_score(personal_token string, owner string, repo string) float64 {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: personal_token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	contributors_count, _, err := client.Repositories.ListContributorsStats(ctx, owner, repo)
	if err != nil {
		if _, ok := err.(*github.AcceptedError); ok {
			log.Println("The information required is not yet ready and was scheduled on GitHub side. Please try again later.")
			return 0
		} else {
			log.Println("Error:", err)
			return 0
		}
	}
	sugar_logger.Debugf("Number of contributors: %d", len(contributors_count))
    score, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", float64(len(contributors_count)) / float64(100)), 64)
    return score
}