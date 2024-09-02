//lint:file-ignore U1000 Ignore all unused code
package main

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc]
	if originalColor == color {
		return image
	}
	image[sr][sc] = color
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, direction := range directions {
		newSr, newSc := sr+direction[0], sc+direction[1]
		if 0 <= newSr && newSr < len(image) && 0 <= newSc && newSc < len(image[0]) && image[newSr][newSc] == originalColor {
			floodFill(image, newSr, newSc, color)
		}
	}
	return image
}
