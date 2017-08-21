package main

import (
	"fmt"
	"strconv"
)

func main() {
	var ampm, military string

	fmt.Scan(&ampm)

	suffix := ampm[8:10]
	hours, _ := strconv.Atoi(ampm[0:2])

	switch suffix {
	case "PM":
		if hours != 12 {
			hours += 12
		}
	case "AM":
		if hours == 12 {
			hours = 0
		}
	}

	military = strconv.Itoa(hours) + ampm[2:8]

	if hours < 10 {
		military = "0" + military
	}

	fmt.Print(military)
}
