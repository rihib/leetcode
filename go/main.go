package main

import "fmt"

func main() {
	image := floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2)
	for _, v := range image {
		for _, b := range v {
			fmt.Println(b)
		}
	}
}
