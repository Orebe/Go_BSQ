package main

import "os"

// Function for checking error.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	myMap, col, row := openFileToMap(os.Args[1])
	itab := findObstacle(myMap, col, row)
	myMap = findSquare(myMap, itab, col, row)
	displayMap(myMap)
}
