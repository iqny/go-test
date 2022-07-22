package main

import (
	"fmt"
	"go-zh/ch4/github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	now := time.Now()
	prMonth := now.AddDate(0, -1, 0)
	prYear := now.AddDate(-1, 0, 0)

	var mIssues []*github.Issue
	var yIssues []*github.Issue
	var overIssues []*github.Issue
	for _, item := range result.Items {
		created := item.CreatedAt.Unix()
		if created > prMonth.Unix() {
			mIssues = append(mIssues, item)
			continue
		}
		if created < prMonth.Unix() && created > prYear.Unix() {
			yIssues = append(yIssues, item)
			continue
		}
		overIssues = append(overIssues, item)
	}
	fmt.Println("issues(不到一个月):")
	for _, item := range mIssues {
		fmt.Printf("#%-5d %9.9s %.55s 时间:%s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}

	fmt.Println("issues(不到一年):")
	for _, item := range yIssues {
		fmt.Printf("#%-5d %9.9s %.55s 时间:%s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	fmt.Println("issues(超过一年):")
	for _, item := range overIssues {
		fmt.Printf("#%-5d %9.9s %.55s 时间:%s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}
