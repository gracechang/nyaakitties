package follow

import (
	"gopkg.in/ahmdrz/goinsta.v2"
	"log"
	"math/rand"
	"throttle"
	"time"
)

// todo: struct with login

func follow(i goinsta.Item, insta *goinsta.Instagram, throttleTrack *throttle.ThrottleTracker) {
	if len(i.Images.Versions) != 0 {
		username := i.User.Username
		log.Println("Trying to follow: ", username)
		user, err := insta.Profiles.ByName(username)
		if err != nil {
			log.Println("Failed to gather profile for ", username)
			log.Println(err)
			throttle.HandleThrottle(throttleTrack)
			return
		}
		followErr := user.Follow()
		if followErr != nil {
			log.Println("Failed to follow ", username)
			log.Println(followErr)
			throttle.HandleThrottle(throttleTrack)
			return
		}
		log.Println("Followed: ", username)
	}
}

func FollowerService(hashtagName string, throttle *throttle.ThrottleTracker, insta *goinsta.Instagram) {
	hashtagInsta := insta.NewHashtag(hashtagName)
	for hashtagInsta.Next() {
		sectionCount := len(hashtagInsta.Sections)
		for i := 0; i < sectionCount; i++ {
			log.Println("asdfadasdf")
			processFollowerSection(hashtagInsta, i, throttle, insta)
		}
	}
}

func processFollowerSection(hashtagInsta *goinsta.Hashtag, sectionNumber int, throttleTrack *throttle.ThrottleTracker, insta *goinsta.Instagram) {
	if len(hashtagInsta.Sections) > 0 {
		for _, media := range hashtagInsta.Sections[sectionNumber].LayoutContent.Medias {
			if throttle.NeedToThrottle(throttleTrack)   {
				follow(media.Item, insta, throttleTrack)
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			}
		}
	}
}
