package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)
	p()

	then := time.Date(2000, 1, 2, 10, 30, 59, 651387237, time.UTC)
	p(then)
	p()

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p()

	p(then.Weekday())
	p()

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))
	p()

	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	p()

	p(then.Add(diff))
	p(then.Add(-diff))
}
