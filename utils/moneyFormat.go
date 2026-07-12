package utils

import "strconv"

func MoneyFormat(total *int) string {
	totalStr := strconv.Itoa(*total)
	format := ""
	switch len(totalStr) {
	case 4:
		format = totalStr[:1] + "." + totalStr[1:]
	case 5:
		format = totalStr[:2] + "." + totalStr[2:]
	case 6:
		format = totalStr[:3] + "." + totalStr[3:]
	case 7:
		format = totalStr[:1] + "." + totalStr[1:4] + "." + totalStr[4:]
	}
	return format
}
