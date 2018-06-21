package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Function for opening a file and return a array of string in myMap.
func openFileToMap(f string) ([]string, int, int) {
	var myMap []string

	file, err := os.Open(f)
	check(err)
	defer file.Close()
	var col, row int
	scanner := bufio.NewScanner(file)
	b := false
	for scanner.Scan() {
		if b == false {
			str := scanner.Text()
			tab := strings.Split(str, " ")
			col, _ = strconv.Atoi(tab[0])
			row, _ = strconv.Atoi(tab[1])
			b = true
		} else {
			myMap = append(myMap, scanner.Text())
		}
	}

	err = scanner.Err()
	check(err)

	return myMap, col, row
}
