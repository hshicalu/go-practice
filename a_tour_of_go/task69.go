
package main

import "fmt"

type Person struct {
    Name string
    Age int
    }

func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
    }

func main() {
    a := Person{"Hicalu", 24}
    z := Person{"Alex", 25}

    fmt.Println(a, z)
    }
