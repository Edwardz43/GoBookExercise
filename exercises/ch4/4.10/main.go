package main

import (
	"GoBook/book/ch4/github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	format := "#%-5d %9.9s %.55s\n"
	now := time.Now()

	arr := make(map[string][]*github.Issue)

	pastDay := make([]*github.Issue, 0)
	pastMonth := make([]*github.Issue, 0)
	pastYear := make([]*github.Issue, 0)
	overYear := make([]*github.Issue, 0)

	day := now.AddDate(0, 0, -1)
	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)

	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(day):
			pastDay = append(pastDay, item)
		case item.CreatedAt.After(month):
			pastMonth = append(pastMonth, item)
		case item.CreatedAt.After(year):
			pastYear = append(pastYear, item)
		case item.CreatedAt.Before(year):
			overYear = append(overYear, item)
		}
	}

	arr["pastDay"] = pastDay
	arr["pastMonth"] = pastMonth
	arr["pastYear"] = pastYear
	arr["overYear"] = overYear

	for k, v := range arr {
		if len(v) > 0 {
			fmt.Printf("\n%s:\n", k)
			for _, item := range v {
				fmt.Printf(format, item.Number, item.User.Login, item.Title)
			}
		}
	}
}
