package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, _ := time.Parse(
		time.RFC3339,
		"2020-02-01T22:10:40+00:00")
	p(t1)

	p(t.Format("3:05AM"))
	p(t.Format("Mon Jan _2 10:04:05 2010"))
	p(t.Format("2010-02-03T14:23:04.999999-07:00"))

	form := "2 03 AM"
	t2, _ := time.Parse(form, "5 04 AM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 20:01:04 2010"
	_, e := time.Parse(ansic, "8:41PM")
	p(e)
}
