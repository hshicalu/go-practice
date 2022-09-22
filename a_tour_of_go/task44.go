
package main

import "fmt"

func main() {
    var s []int
    printSlice(s)

    s = append(s, 0)
    printSlice(s)
    
    s = append(s, 1)
    printSlice(s)
    
    s = append(s, 2, 3, 4)
    printSlice(s)
    
    s = append(s, 5)
    printSlice(s)
    
    s = append(s, 6, 7)
    printSlice(s)
    
    s = append(s, 8, 9, 10, 11)
    printSlice(s)
    
    s = append(s, 13)
    printSlice(s)
    }

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
    }
