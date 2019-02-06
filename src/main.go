package main

import (
	"fmt"
	"follow"
	"gopkg.in/ahmdrz/goinsta.v2"
	"like"
	"log"
	"os"
	"strconv"
	"throttle"
	"time"
)

func logIn() (*goinsta.Instagram) {
	insta, err := goinsta.Import("~/.goinsta")
	if err != nil {
		log.Println(err)
		log.Println("Trying environment variables...")
		username := os.Getenv("USERNAME")
		pass := os.Getenv("PASSWORD")
		if username == "" || pass == "" {
			log.Fatal("Unable to get credentials.")
			os.Exit(1)
		}
		insta = goinsta.New(username, pass)
	}
	if loginErr := insta.Login(); loginErr != nil {
		fmt.Println(loginErr)
		os.Exit(1)
	}
	insta.Export("~/.goinsta")
	log.Println("Successfully logged in.")
	return insta
}

func main() {
	// todo: take in a list of hashtags
	followThrottleCounter := throttle.ThrottleTracker{RejectCount: 0, LastReject: 0}
	followCount, _ := strconv.Atoi(os.Getenv("FOLLOW_COUNT"))
	for i := 0; i < followCount; i++ {
		go follow.FollowerService("cats", &followThrottleCounter, logIn()) // todo: add waitgroup
	}
	likeThrottleCounter := throttle.ThrottleTracker{RejectCount: 0, LastReject: 0}
	likeCount, _ := strconv.Atoi(os.Getenv("LIKE_COUNT"))
	for i := 0; i < likeCount; i++ {
		go like.LikeService("cats", &likeThrottleCounter, logIn()) // todo: add waitgroup
	}
	time.Sleep(10 * time.Hour) // todo: make this not dumb
}
