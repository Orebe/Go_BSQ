package main

// Function who find obstacle in map.
func findObstacle(mapp []string, col int, row int) [][]int {
	itab := make([][]int, col)
	for i := 0; i < col; i++ {
		itab[i] = make([]int, row)
		for y := 0; y < row; y++ {
			if i == 0 || y == 0 {
				if mapp[i][y] == 'o' {
					itab[i][y] = 1
				} else {
					itab[i][y] = 0
				}
			} else {
				itab[i][y] = itab[i][y-1] + itab[i-1][y] - itab[i-1][y-1]
				if mapp[i][y] == 'o' {
					itab[i][y] = itab[i][y] + 1
				}
			}
		}
	}
	return (itab)
}

// Function who find the biggest square possible in map.
func findSquare(mapp []string, itab [][]int, col int, row int) []string {
	size, sx, sy := 0, 0, 0
	for i := 0; i < col; i++ {
		for y := 0; y < row && y+size < row; y++ {
			for i+size < col && y+size < row && itab[i+size][y+size]-itab[i+size][y]-itab[i][y+size]+itab[i][y] <= 0 {
				sx = i
				sy = y
				size++
			}
		}
	}
	for i := sx; i < sx+size; i++ {
		for y := sy; y < sy+size; y++ {
			mapp[i] = mapp[i][:y] + "x" + mapp[i][y+1:]
		}
	}
	return mapp
}
