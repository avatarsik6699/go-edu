package repeater

import "strings"

func Repeater(str string, count int) string {
	var res strings.Builder

	for range count {
		res.WriteString(str)
	}

	return res.String()
}
