//lint:file-ignore U1000 Ignore all unused code
package main

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]strings.Builder, numRows)
	cycleLen := 2*numRows - 2
	for i, r := range s {
		rowNum := i % cycleLen
		if rowNum >= numRows {
			rowNum = cycleLen - rowNum
		}
		rows[rowNum].WriteRune(r)
	}

	var oneline strings.Builder
	for _, row := range rows {
		oneline.WriteString(row.String())
	}
	return oneline.String()
}
