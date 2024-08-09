//lint:file-ignore U1000 Ignore all unused code
package zigzagconversion

/*
より読みやすくなるようにリファクタしました。
*/
func convert_step2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([][]rune, numRows)
	cycleLen := 2*numRows - 2
	for i, r := range s {
		rowNum := i % cycleLen
		if rowNum >= numRows {
			rowNum = cycleLen - rowNum
		}
		rows[rowNum] = append(rows[rowNum], r)
	}

	var oneline []rune
	for _, row := range rows {
		oneline = append(oneline, row...)
	}
	return string(oneline)
}
