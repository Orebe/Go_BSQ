package main

import (
    "fmt"
)

// Function for displaying Map.
func displayMap(myMap []string) {
    for _, value := range myMap {
        fmt.Println(value)
    }
}