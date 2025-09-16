package main

import (
	"fmt"
	"math/rand"
)

func generateRandomSlice(ch chan int) {
	arr := make([]int, 10)

	for idx := range arr {
		arr[idx] = rand.Intn(100)
	}

	for _, item := range arr {
		ch <- item
	}

	fmt.Printf("Отправлено: %d\n", arr)
	close(ch)
}

func square(ch chan int, outCh chan int)  {
	for num := range ch {
		outCh <- num * num
	}
	close(outCh)
}

func main()  {
	ch := make(chan int)
	outCh := make(chan int)

	go generateRandomSlice(ch)
	go square(ch, outCh)

	for elem := range outCh {
		fmt.Println(elem)
	}
}
