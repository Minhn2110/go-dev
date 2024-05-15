package main

import (
	"fmt"
	"time"
)

func main() {

	ExampleTimeZone()
}

func ExampleTimeZone() {
	//s := "2024-05-03T10:34:36-09:00"
	//
	//t, err := time.Parse(time.RFC3339, s)
	//
	//if err != nil {
	//	fmt.Println("The string doesn't match a time.Time value:", err)
	//} else {
	//	fmt.Println("t", t.UTC())
	//	fmt.Println("t", t.Local())
	//	name, zoneOrOffset := t.Zone()
	//	fmt.Println("The timezone is:", name)
	//	fmt.Println("The timezone is:", zoneOrOffset)
	//}

	t := time.Now()        // get the current time
	weekday := t.Weekday() // get the day of the week

	fmt.Println("Today is:", weekday)
}
