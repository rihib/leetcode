//lint:file-ignore U1000 Ignore all unused code
package main

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc]

	if originalColor == color {
		return image
	}
	image[sr][sc] = color

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range directions {
		nSr, nSc := sr+d[0], sc+d[1]
		if 0 <= nSr && nSr < len(image) &&
			0 <= nSc && nSc < len(image[0]) &&
			image[nSr][nSc] == originalColor {
			image = floodFill(image, nSr, nSc, color)
		}
	}

	return image
}
