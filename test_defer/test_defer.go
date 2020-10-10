package main

import "fmt"

func test() (z int) {
	defer func() {
		fmt.Println("defer:", z)
		z += 100
	}()
	return 100
}

func main0() {
	fmt.Println("test:", test())

	const (
		i = 1 << iota
		j
		k
		l
	)

	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}

func main() {
	fmt.Println("main start")
	ch := make(chan string)
	go func() {
		val := <-ch // 出 chan
		fmt.Println(val)
	}()
	ch <- "a" // 入 chan
	fmt.Println("main end")
}
