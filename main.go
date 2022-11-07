package main

import (
	"fmt"
	"time"
)

func main() {
	st := time.Now()

	fmt.Printf("Start time : %s\n", st.Format(time.RFC3339))

	// go func() {
	// 	for {
	// 		for _, r := range `-\|/` {
	// 			fmt.Printf("\r%c", r)
	// 			time.Sleep(time.Microsecond * 10000)
	// 		}
	// 	}
	// }()
	// go calculatorSonik(1000)
	// go calculatorSonik(2000)
	// time.Sleep(100 * time.Millisecond)
	// fmt.Printf("end time : %s\n", time.Since(st))

	result1 := make(chan int)
	result2 := make(chan int)

	go calculatorChan(1000, result1)
	go calculatorChan(2000, result2)

	fmt.Println(<-result1)
	fmt.Println(<-result2)

	fmt.Printf("end time : %s\n", time.Since(st))
}

func calculatorSonik(n int) {
	t := time.Now()
	result := 0

	for i := 0; i < n; i++ {
		result += i * 2
		time.Sleep(time.Microsecond * 3)
	}

	fmt.Printf("Result of Calc: %d; Spended time : %s\n", result, time.Since(t))
}

func calculatorChan(n int, res chan int) {
	t := time.Now()
	result := 0

	for i := 0; i < n; i++ {
		result += i * 2
		time.Sleep(time.Microsecond * 3)
	}

	fmt.Printf("Spended time : %s\n", time.Since(t))
	res <- result
}
