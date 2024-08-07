//lint:file-ignore U1000 Ignore all unused code
package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([][]rune, numRows)
	for i, r := range s {
		offset := i % (numRows*2 - 2)
		var rowNum int
		if offset < numRows {
			rowNum = offset
		} else {
			rowNum = 2*numRows - offset - 2
		}
		rows[rowNum] = append(rows[rowNum], r)
	}

	var oneline []rune
	for _, row := range rows {
		oneline = append(oneline, row...)
	}
	return string(oneline)
}
