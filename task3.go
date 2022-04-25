
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    rand.Seed(2022)
    fmt.Println("Num is ", rand.Intn(10))
    }
