package format

import (
	"strconv"
	"strings"
)

func Result(expression string, result float64) string {
	rawStr := strconv.FormatFloat(result, 'f', 2, 64)

	parts := strings.Split(rawStr, ".")
	body := parts[0]
	tails := "." + parts[1]

	n := len(body)
	var res strings.Builder

	res.Grow(n + (n / 3) + len(tails))

	for i := range n {
		if i > 0 && (n-i)%3 == 0 {
			res.WriteByte('.')
		}
		res.WriteByte(body[i])
	}

	res.WriteString(tails)

	return res.String()
}
