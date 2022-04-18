package main

import (
	"fmt"
	"time"
)

var p = fmt.Println

func main() {
	now := time.Now()
	p(now)
	p(now.Unix())
	then := time.Date(2001, 8, 16, 20, 34, 58, 0, time.UTC)

	p(then)

	p("Year", now.Year())
	p("Month", now.Month())
	p("Day", now.Day())
	p("Hour", now.Hour())
	p("Second", now.Second())
	p("Nanosecond", now.Nanosecond())
	p("Location", now.Location())
	p("Weekday", now.Weekday())

	p("Before", then.Before(now))
	p("After", then.After(now))
	p("Equal", then.Equal(now))

	diff := now.Sub(then)
	p("diff", diff)

	p(diff.Hours())
}
