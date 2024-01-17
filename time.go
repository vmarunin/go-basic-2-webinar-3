package main

import (
	"fmt"
	"time"
)

func ExampleTime() {
	ExampleTimeSleep()
	// ExampleTimeMeasure()
	// ExampleTimeParse()
	// ExampleTimeFriday13()
}

func ExampleTimeSleep() {
	fmt.Println("Before sleep")
	time.Sleep(5 * time.Second)
	fmt.Println("After sleep")
}

func ExampleTimeMeasure() {
	now := time.Now()
	coiunter := 0
	for i := 0; i < 1_000_000_000; i++ {
		coiunter++
	}
	duration := time.Since(now)
	fmt.Println(duration.String())
}

/*
01/02 03:04:05PM '06 -0700
Year: "2006" "06"
Month: "Jan" "January" "01" "1"
Day of the week: "Mon" "Monday"
Day of the month: "2" "_2" "02"
Day of the year: "__2" "002"
Hour: "15" "3" "03" (PM or AM)
Minute: "4" "04"
Second: "5" "05"
AM/PM mark: "PM"
*/
func ExampleTimeParse() {
	loc, _ := time.LoadLocation("Europe/Berlin")

	// This will look for the name CEST in the Europe/Berlin time zone.
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	today, _ := time.ParseInLocation(longForm, "Oct 28, 2023 at 5:02pm (CEST)", loc)
	fmt.Println(today)
	tomorrow := today.Add(time.Hour * 24)
	tomorrowRight := today.AddDate(0, 0, 1)
	fmt.Println(tomorrow)
	fmt.Println(tomorrowRight)
}

// Доказать, что 13 число чаще выпадает на пятницу, чем на другие дни недели
func ExampleTimeFriday13() {
	day := time.Date(2000, time.January, 13, 12, 0, 0, 0, time.UTC)
	endDay := time.Date(2400, time.January, 1, 12, 0, 0, 0, time.UTC)
	counter := map[string]int{}
	for day.Before(endDay) {
		counter[day.Weekday().String()]++
		day = day.AddDate(0, 1, 0)
	}
	fmt.Println(counter)
}

// Тикер и таймер будут потом
