package main

import "fmt"

func main() {
	items := []string{"a", "b", "c"}
	var cnt int
	for i := 0; i < len(items); i++ {
		if items[i] == "b" {
			if cnt <= 2 {
				items = append(items, items[i])
				cnt += 1
				fmt.Printf("Item is [%s] in IF.\n", items[i])
				continue
			}
		}
		fmt.Println(items)
		fmt.Printf("Item is [%s]\n", items[i])
	}
}
