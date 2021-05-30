//copyright

// auto forces null blocks - beta
// also helps with POW equation

package main

import (
	"time"
)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}


