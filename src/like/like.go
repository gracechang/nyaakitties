package like

import (
	"gopkg.in/ahmdrz/goinsta.v2"
	"log"
	"throttle"
)

type LikeTracker struct {
	LikeCount int64
}

// todo: struct with login
// todo: clean up

func like(likeItem goinsta.Item, insta *goinsta.Instagram, likeTracker *LikeTracker, throttler *throttle.ThrottleTracker) {
	if len(likeItem.Images.Versions) != 0 {
		mediaId := likeItem.ID
		media, _ := insta.GetMedia(mediaId)
		if len(media.Items) > 0 {
			if throttle.NeedToThrottleLike(throttler) {
				likeErr := media.Items[0].Like()
				log.Println("Liking:", media.Items[0].User.Username, media.Items[0].HasLiked)
				if likeErr != nil {
					throttler.LastReject = throttle.CurrentMillis()
					throttler.RejectCount += 1
					log.Fatal("Failed to like photo for ", media.Items[0].User.Username)
					log.Fatal(likeErr) // todo: add logging for throttler
				} else {
					likeTracker.LikeCount += 1
					if likeTracker.LikeCount % 100 == 0 {
						log.Println("Like Count Update:", likeTracker.LikeCount) // todo: make better log
					}
				}

			}
		}
	}
}

func LikeService(hashtagName string, throttler *throttle.ThrottleTracker, insta *goinsta.Instagram) {
	likeTracker := LikeTracker{LikeCount:0}
	h := insta.NewHashtag(hashtagName)
	for h.Next() {
		for i := range h.Sections {
			for _, i := range h.Sections[i].LayoutContent.Medias {
				like(i.Item, insta, &likeTracker, throttler)
			}
		}
	}
}