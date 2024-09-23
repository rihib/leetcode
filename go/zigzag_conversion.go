//lint:file-ignore U1000 Ignore all unused code
package main

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]strings.Builder, numRows)
	cycleLength := 2*numRows - 2
	for i, r := range s {
		rowIndex := i % cycleLength
		if rowIndex >= numRows {
			rowIndex = cycleLength - rowIndex
		}
		rows[rowIndex].WriteRune(r)
	}

	var oneline strings.Builder
	for _, row := range rows {
		oneline.WriteString(row.String())
	}
	return oneline.String()
}

func convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var oneline strings.Builder
	cycleLength := 2*numRows - 2
	for rowIndex := 0; rowIndex < numRows; rowIndex++ {
		for i := rowIndex; i < len(s); i += cycleLength {
			oneline.WriteByte(s[i])
			if rowIndex == 0 || rowIndex == numRows-1 {
				continue
			}
			diagIndex := i + cycleLength - rowIndex*2
			if diagIndex < len(s) {
				oneline.WriteByte(s[diagIndex])
			}
		}
	}
	return oneline.String()
}
