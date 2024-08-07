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

func convert_anothersolution(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var oneline strings.Builder
	cycleLen := 2*numRows - 2
	for rowNum := 0; rowNum < numRows; rowNum++ {
		for i := rowNum; i < len(s); i += cycleLen {
			oneline.WriteByte(s[i])
			if 0 < rowNum && rowNum < numRows-1 {
				diagIdx := i + cycleLen - rowNum*2
				if diagIdx < len(s) {
					oneline.WriteByte(s[diagIdx])
				}
			}
		}
	}
	return oneline.String()
}
