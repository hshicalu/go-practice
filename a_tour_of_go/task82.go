// Check "for" behavior
package main

import "fmt"

func main() {
	items := []string{"a", "b", "c"}
	var cnt int
	// pattern-1
	// for i, item := range items {
	// 	fmt.Printf("Item is : [%s]\n", item)
	// 	if item == "b" {
	// 		if cnt <= 2 {
	// 			items = append(items, item)
	// 			cnt += 1
	// 			continue
	// 		}
	// 	}
	// 	fmt.Printf("Loop is : %d\n", i+1)
	// 	fmt.Printf("Items is : %s\n", items)
	// }
	// Item is : [a]
	// Loop is : 1
	// Items is : [a b c]
	// Item is : [b]
	// Item is : [c]
	// Loop is : 3
	// Items is : [a b c b]

	// pattern-2
	for i := 0; i < len(items); i++ {
		fmt.Printf("Item is : [%s]\n", items[i])
		if items[i] == "b" {
			if cnt <= 2 {
				items = append(items, items[i])
				cnt += 1
				continue
			}
		}
		fmt.Printf("Loop is : %d\n", i+1)
		fmt.Printf("Items is : %s\n", items)
	}
	// 	Item is : [a]
	// Loop is : 1
	// Items is : [a b c]
	// Item is : [b]
	// Item is : [c]
	// Loop is : 3
	// Items is : [a b c b]
	// Item is : [b]
	// Item is : [b]
	// Item is : [b]
	// Loop is : 6
	// Items is : [a b c b b b]
}
