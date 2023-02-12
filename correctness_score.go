package main

import (
	"context"
	"fmt"
	"log"
    "strconv"
	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

func correctness_score(personal_token string, owner string, repo string) float64 {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: personal_token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	weekly_activity, _, err := client.Repositories.ListCodeFrequency(ctx, owner, repo)
	if err != nil {
		if _, ok := err.(*github.AcceptedError); ok {
			log.Println("scheduled on GitHub side")
			return 0
		} else {
			log.Println(err)
			return 0
		}
	}
	if len(weekly_activity) == 0 {
		return 1.0
	} else {
		deletions := *(weekly_activity[0].Deletions)
		if *(weekly_activity[0].Deletions) < 0 {
			deletions = -1 * (*(weekly_activity[0].Deletions))
		}
		sugar_logger.Debugf("Number of deletions: %d", deletions)
		additions := *(weekly_activity[0].Additions)
		sugar_logger.Debugf("Number of additions: %d", *(weekly_activity[0].Additions))
		score, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", 1 - float64(deletions) / float64(additions)), 64)
		if score < 0 {
			score = -1 * score
		}
		return score
	}
}