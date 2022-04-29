
package main

import "fmt"

type Vertex struct {
    Lat, Long float64
    }

var m = map[string]Vertex {
    "Bell Labs": Vertex {
        40.68433, -74.39967, 
    }, 
    "Google": Vertex {
        37.42202, -122.08408, 
    }, 
}

func main() {
    // fmt.Println(m)
    // fmt.Println(m["Bell Labs"])
    // fmt.Println(m["Bell Labs"].Lat)
    fmt.Println(m["Gogle"])
    fmt.Println(m["Gogle"].Lat)
    // fmt.Println(m["Google"].Long)
    }
