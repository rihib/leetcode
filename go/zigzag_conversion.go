//lint:file-ignore U1000 Ignore all unused code
package main

import "strings"

// ジグザグの行ごとに作成していき、最後に１つにする
// sの先頭から見ていくので、その文字の行番号さえわかればその行に追加していくだけ
func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	rows := make([]strings.Builder, numRows)
	blockSize := numRows*2 - 2
	for i, r := range s {
		offset := i % blockSize
		var rowIndex int
		if offset < numRows {
			rowIndex = offset
		} else {
			rowIndex = blockSize - offset
		}
		rows[rowIndex].WriteRune(r)
	}
	var converted strings.Builder
	for _, row := range rows {
		converted.WriteString(row.String())
	}
	return converted.String()
}
