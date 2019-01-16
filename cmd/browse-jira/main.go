package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	client "git.ndumas.com/ndumas/slayer/client"
	jira "git.ndumas.com/ndumas/slayer/jira"
)

func main() {
	var username = flag.String("user", "", "jira username")
	var token = flag.String("token", "", "jira token")

	var board = flag.Int("board", 0, "filter id")
	var filter = flag.Int("filter", 0, "board id")

	flag.Parse()

	log.Printf("username: %s", username)
	log.Printf("token: %s", token)

	if *board != 0 && *filter != 0 {
		log.Fatal("please select ONLY board or filter, not both")
	}

	auth := client.AuthOptions{User: *username, Token: *token}

	js := jira.New(auth, "", time.Second*30)
	fmt.Printf("%+v", js)
}