package repository

import (
	"strconv"
	"strings"
)

func getValuesString(argCount, rowsCount, startWith int) string {
	if argCount < 1 {
		return ""
	}
	var q strings.Builder
	for i := 0; i < rowsCount; i++ {
		q.WriteString("(")
		for j := 0; j < argCount; j++ {
			if j > 0 {
				q.WriteString(", ")
			}
			q.WriteString("$")
			q.WriteString(strconv.Itoa(startWith + i*argCount + j))
		}
		q.WriteString("),")
	}

	return strings.TrimSuffix(q.String(), ",")
}
