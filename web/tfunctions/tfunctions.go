package tfunctions

import "strconv"

func Percent(f float64) string {
	return strconv.Itoa(int(f*100)) + "%"
}
