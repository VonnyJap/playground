package main

import "fmt"

func main() {
	fmt.Println(minHeightShelves([][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, 4))
}

type Shelf struct {
	Height int
	Width  int
}

func minHeightShelves(books [][]int, shelfWidth int) int {

	if len(books) == 0 {
		return 0
	}
	if len(books) == 1 {
		return books[0][1]
	}
	result := []Shelf{}
	result = append(result, Shelf{Height: books[0][1], Width: books[0][0]})
	heightSofar := books[0][1]
	if result[0].Width+books[1][0] <= shelfWidth {
		result = append(result, Shelf{Height: max(result[0].Height, books[1][1]), Width: result[0].Width + books[1][0]})
		heightSofar = max(result[0].Height, books[1][1])
	} else {
		result = append(result, Shelf{Height: result[0].Height + books[1][1], Width: books[1][0]})
		heightSofar = books[1][1]
	}

	// need global height and local height
	for i := 2; i < len(books); i++ {
		// need to update this heightSofar
		height := max(heightSofar, books[i][1])
		width := result[i-1].Width + books[i][0]
		if width > shelfWidth {
			width = books[i][0]
			heightSofar = books[i][1]
		}
		// yes or no is in loop?
		fmt.Println("height, width: ", height, width)
		loopWidth := books[i][0]
		loopHeight := books[i][1]
		for j := i - 1; j > 0; j-- {
			loopWidth += books[j][0]
			if loopWidth > shelfWidth {
				break
			}
			loopHeight = max(loopHeight, books[j][1])
			fmt.Println("loopHeight, loopWidth: ", loopHeight, loopWidth)
			if result[j-1].Height+loopHeight < height {
				height = result[j-1].Height + loopHeight
				width = loopWidth
			}
		}
		result = append(result, Shelf{Height: height, Width: width})
		fmt.Println(result)
	}

	return result[len(result)-1].Height
}

// public int minHeightShelves(int[][] books, int shelfWidth) {
// 	int memo[] = new int[books.length + 1];

// 	for (int index = books.length - 1; index >= 0; index--) {
// 			int ans = Integer.MAX_VALUE;
// 			int height = 0;
// 			int width = 0;

// 			for (int i = index; i < books.length; i++) {
// 					width += books[i][0];

// 					if (width > shelfWidth) {
// 							break;
// 					}
// 					height = Math.max(height, books[i][1]);
// 					ans = Math.min(ans, height + memo[i + 1]);
// 			}
// 			memo[index] = ans;
// 	}
// 	return memo[0];
// }
func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}

func min(first, second int) int {
	if first < second {
		return first
	}
	return second
}
