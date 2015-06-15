package main

import (
	"fmt"
	"github.com/google/go-github/github"
	"github.com/raguay/goAlfred"
	"os"
	"strings"
)

func main() {

	if len(os.Args) > 2 {
		// Load up creds if present
		gh := github.NewClient(nil)
		switch os.Args[1] {
		case "login":
			// TODO: save creds for authenticating to google
		case "logout":
			// TODO: delete creds
		case "keys-for":
			// save username into cache, ++ing usage count
			var keys []string
			page := 0
			for true {
				results, response, err := gh.Users.ListKeys(os.Args[2], &github.ListOptions{Page: page, PerPage: 500})
				if err != nil {
					panic(err)
				}
				page = response.NextPage
				for i := 0; i < len(results); i++ {
					keys = append(keys, *results[i].Key)
				}
				if response.NextPage == response.LastPage {
					break
				}
			}
			fmt.Print(strings.Join(keys, "\n") + "\n")

		case "find-user":
			results, _, err := gh.Search.Users(os.Args[2], &github.SearchOptions{})
			if err != nil {
				panic(err)
			}
			for i := 0; i < len(results.Users); i++ {
				// if user is in cache, use value as priority
				user := results.Users[i]
				id := *user.ID
				login := *user.Login
				var name string
				if user.Name != nil {
					name = *user.Name
				} else {
					name = *user.Login
				}
				goAlfred.AddResult(
					fmt.Sprintf("%d", id), // uid
					login, // arg string
					login, // title
					"Copy SSH keys for "+name+" to clipboard", // subtitle
					"icon.png", // icon
					"yes",      // valid
					"",         // auto
					"",         // rtype
					0,          // priority
				)
			}
			fmt.Print(goAlfred.ToXML())
		}
	}
}
