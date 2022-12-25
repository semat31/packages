package main

import (
	"fmt"
	"strconv"
	"time"
)

type CWeek struct {
	WeekISO string

	//WeekStart time.Time
	//WeekEnd time.Time
}

func CWeekNew(t time.Time) (cw CWeek) {
	year, week := t.ISOWeek()
	cw.WeekISO = strconv.Itoa(year) + "W" + strconv.Itoa(week)

	return
}

func WeekStart(year, week int) time.Time {
	// Start from the middle of the year:
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)

	// Roll back to Monday:
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	// Difference in weeks:
	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)

	return t
}

func WeekRange(year, week int) (start, end time.Time) {
	start = WeekStart(year, week)
	end = start.AddDate(0, 0, 6)
	return
}

func main() {
	var t_day = "25.12.2022"

	var date_layout = "02.01.2006"

	t, err := time.Parse(date_layout, t_day)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Извлечение единиц из: %v\n", t)
	dOfMonth := t.Day()
	weekDay := t.Weekday()
	month := t.Month()
	fmt.Printf("%d-й день из %v это %v\n", dOfMonth, month, weekDay)

	cWeek := CWeekNew(t)
	fmt.Println(cWeek.WeekISO)

	fmt.Println(WeekRange(2022, 51))
	fmt.Println(WeekRange(2018, 2))
	fmt.Println(WeekRange(2019, 1))
	fmt.Println(WeekRange(2019, 2))
}
