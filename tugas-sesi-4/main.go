package main

import (
	"fmt"
	"sync"
)

func main() {

	numbers := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	wg := sync.WaitGroup{}

	wg.Add(len(numbers))
	for _, v := range numbers {
		go func(num int) {
			fmt.Println(num)
			wg.Done()
		}(v)
	}
	wg.Wait()

}
