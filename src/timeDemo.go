package src

import "time"
import "fmt"

func DateDemo() time.Time {
	fmt.Println(time.Now())
	return time.Now()
}

func WeekDayDemo(time2 time.Time) time.Weekday {
	fmt.Println(time2.Weekday())
	return time2.Weekday()
}
