package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCurrentQuarter(t *testing.T) {
	cases := []struct {
		month   string
		quarter int
	}{
		{month: "01", quarter: 1},
		{month: "02", quarter: 1},
		{month: "03", quarter: 1},
		{month: "04", quarter: 2},
		{month: "05", quarter: 2},
		{month: "06", quarter: 2},
		{month: "07", quarter: 3},
		{month: "08", quarter: 3},
		{month: "09", quarter: 3},
		{month: "10", quarter: 4},
		{month: "11", quarter: 4},
		{month: "12", quarter: 4},
	}

	//TODO Реализовать Календарь

	for _, test := range cases {
		parsed, _ := time.Parse("2006-01-02", fmt.Sprintf("2015-%s-15", test.month))
		calendar := NewCalendar(parsed)
		actual := calendar.CurrentQuarter()
		if actual != test.quarter {
			t.Error("Month:", test.month,
				"Expected Quarter:", test.quarter,
				"Actual Quarter:", actual)
		}
	}
}

//type NewCalendar time.Time

type calendar struct {
	CurrentTime time.Time
}

func NewCalendar(CurrentTime time.Time) *calendar  {
	return &calendar{CurrentTime:CurrentTime}
}

func (c *calendar) CurrentQuarter() int  {
	month :=c.CurrentTime.Month()
	switch month.String() {
	case "January", "February","March": return 1
	case "April", "May", "June": return 2
	case "July", "August", "September": return 3
	case "October", "November", "December": return 4
	}
	return 0
}



