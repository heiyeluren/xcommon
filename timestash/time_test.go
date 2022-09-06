package xcommon

import (
	"fmt"
	"log"
	"testing"
	"time"
)

var timeDuration = []int64{1, 2, 3, 4, 5, 12, 24, 30}

func TestGetTimeNow(t *testing.T) {
	time.Sleep(1 * time.Second)
	fmt.Println(GetTimeNow())
	time.Sleep(2 * time.Second)
	fmt.Println(GetTimeNow())
	time.Sleep(3 * time.Second)
	fmt.Println(GetTimeNow())
	time.Sleep(3 * time.Second)
	fmt.Println(GetTimeNow())

	for _, t := range timeDuration {
		lastnow := time.Now().Unix()

		time.Sleep(time.Duration(t) * time.Second)
		now := GetTimeNow()

		log.Println("now: ", now, "lasttime: ", now-lastnow)
	}

}
