package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		year, remainingDays, currentDay, currentMonth int
		date                                          string
	)

	fmt.Scan(&year)

	remainingDays = 255
	currentMonth = 1
	currentDay = 1

	if year == 1918 {
		remainingDays = 224
		currentDay = 14
		currentMonth = 2
	}

	for remainingDays > 0 {
		remainingDays--

		if isLastDayOfMonth(currentDay, currentMonth, year) {
			currentDay = 1
			currentMonth++
		} else {
			currentDay++
		}
	}

	day, _ := padChar(strconv.Itoa(currentDay), 2, '0')
	month, _ := padChar(strconv.Itoa(currentMonth), 2, '0')

	date = fmt.Sprintf("%s.%s.%d", day, month, year)

	fmt.Println(date)
}

func isLeap(year int) bool {
	if year <= 1917 {
		return year%4 == 0
	}

	return year%400 == 0 || (year%100 != 0 && year%4 == 0)
}

func isLastDayOfMonth(day, month, year int) bool {

	switch month {
	case 1:
		if day == 31 {
			return true
		}
	case 2:
		if (day == 29 && isLeap(year)) || (day == 28 && !isLeap(year)) {
			return true
		}
	case 3:
		if day == 31 {
			return true
		}
	case 4:
		if day == 30 {
			return true
		}
	case 5:
		if day == 31 {
			return true
		}
	case 6:
		if day == 30 {
			return true
		}
	case 7:
		if day == 31 {
			return true
		}
	case 8:
		if day == 31 {
			return true
		}
	case 9:
		if day == 30 {
			return true
		}
	case 10:
		if day == 31 {
			return true
		}
	}

	return false
}

// pad left-pads s with spaces, to length n.
// If n is smaller than s, Pad is a no-op.
func pad(s string, n int) (string, error) {
	return padChar(s, n, ' ')
}

// padChar left-pads s with the rune r, to length n.
// If n is smaller than s, PadChar is a no-op.
func padChar(s string, n int, r rune) (string, error) {
	if n < 0 {
		return "", fmt.Errorf("invalid length %d", n)
	}
	if len(s) > n {
		return s, nil
	}
	return strings.Repeat(string(r), n-len(s)) + s, nil
}
