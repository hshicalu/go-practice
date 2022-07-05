package main

import "fmt"

func main() {
	var cnt int
	for {
		cnt += 1
		if cnt%2 == 0 {
			cnt += 2
			continue
		}
		if cnt >= 10 {
			break
		}
		fmt.Println(cnt)
	}

}
