package main

import (
	"fmt"
	"github.com/google/go-github/github"
	"github.com/raguay/goAlfred"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

func record_hit(user string) error {
	cache, err := read_cache()
	if err != nil {
		return err
	}

	cache[user]++

	err = write_cache(cache)
	if err != nil {
		return err
	}
	return nil
}

func read_cache() (map[string]int, error) {
	src, err := ioutil.ReadFile(goAlfred.Cache() + "/cache.yml")
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]int{}, nil
		}
		return nil, err
	}

	var cache map[string]int
	err = yaml.Unmarshal(src, &cache)
	if err != nil {
		return nil, err
	}

	return cache, nil
}

func write_cache(cache map[string]int) error {
	src, err := yaml.Marshal(cache)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(goAlfred.Cache()+"/cache.yml", src, 0644)
	if err != nil {
		return err
	}
	return nil
}

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
			user := os.Args[2]
			var keys []string

			err := record_hit(user)
			if err != nil {
				panic(err)
			}

			page := 0
			for true {
				results, response, err := gh.Users.ListKeys(user, &github.ListOptions{Page: page, PerPage: 500})
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

			cache, err := read_cache()
			if err != nil {
				panic(err)
			}

			for i := 0; i < len(results.Users); i++ {
				user := results.Users[i]
				id := *user.ID
				login := *user.Login
				var name string
				if user.Name != nil {
					name = *user.Name
				} else {
					name = *user.Login
				}
				priority := cache[login]

				goAlfred.AddResult(
					fmt.Sprintf("%d", id), // uid
					login, // arg string
					login, // title
					"Copy SSH keys for "+name+" to clipboard", // subtitle
					"icon.png", // icon
					"yes",      // valid
					"",         // auto
					"",         // rtype
					priority,
				)
			}
			fmt.Print(goAlfred.ToXML())
		}
	}
}
