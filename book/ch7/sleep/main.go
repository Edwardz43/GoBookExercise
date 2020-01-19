package main

import "flag"

import "time"

import "fmt"

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("sleeping for %v\n", *period)
	time.Sleep(*period)
	fmt.Println("aweak")
}
