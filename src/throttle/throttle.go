package throttle

import (
	"log"
	"time"
)

type ThrottleTracker struct {
	RejectCount int32
	LastReject int64
}

func HandleThrottle(throttleTrack *ThrottleTracker) {
	throttleTrack.RejectCount += 1
	throttleTrack.LastReject = CurrentMillis()
	log.Println("THROTTLE => reject count: ", throttleTrack.RejectCount, " last reject: ", throttleTrack.LastReject)
}

func NeedToThrottle(throttleTrack *ThrottleTracker) bool {
	return (throttleTrack.RejectCount < 10 && (CurrentMillis() - throttleTrack.LastReject) > (60000 * 5)) || (CurrentMillis() - throttleTrack.LastReject > 60000 * 60 )
}

func NeedToThrottleLike(throttleTrack *ThrottleTracker) bool { // todo: clean up
	return (throttleTrack.RejectCount < 10 && (CurrentMillis() - throttleTrack.LastReject) > (60000 * 1)) || (CurrentMillis() - throttleTrack.LastReject > 60000 * 15 )
}

func CurrentMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
