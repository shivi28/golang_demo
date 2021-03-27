package main

import (
	"fmt"
	"time"
)

func main() {
	tn := time.Now().UTC()
	fmt.Println(tn)
	year, week := tn.ISOWeek()
	fmt.Println(year, week)

	ts := time.Now().UTC().Unix()
	tn = time.Unix(ts, 0)
	fmt.Println(tn)
	year, week = tn.ISOWeek()
	fmt.Println(year, week)
}
