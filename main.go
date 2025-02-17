package main

import (
	"fmt"
	"testAlg/service"
	"time"
)

func main() {
	fmt.Println("Hi, this is a solution of test task. You can read description of this task at README.md file." +
		"Complexity of my solution is O(n), also i decided to print number of steps at each solution")
	time.Sleep(5 * time.Second)
	fmt.Printf("Number of steps to count: %v\n", service.Solution(400))
	time.Sleep(1 * time.Second)
	fmt.Printf("Number of steps to count: %v\n", service.Solution(200))
	time.Sleep(1 * time.Second)
	fmt.Printf("Number of steps to count: %v\n", service.Solution(0))
	time.Sleep(1 * time.Second)
	fmt.Printf("Number of steps to count: %v\n", service.Solution(-1000))
}
