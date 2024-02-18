//lint:file-ignore U1000 Ignore all unused code
package main

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc]
	image[sr][sc] = color
	if 0 <= sr-1 && image[sr-1][sc] == originalColor {
		image = floodFill(image, sr-1, sc, color)
	}
	if sr+1 < len(image[sr]) && image[sr+1][sc] == originalColor {
		image = floodFill(image, sr+1, sc, color)
	}
	if 0 <= sc-1 && image[sr][sc-1] == originalColor {
		image = floodFill(image, sr, sc-1, color)
	}
	if sc+1 < len(image[sr]) && image[sr][sc+1] == originalColor {
		image = floodFill(image, sr, sc+1, color)
	}
	return image
}
