package service

import (
	"fmt"
	"math/rand"
	"time"
)

var numbers = [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

func Solution(start int) int { // сделаем у ф-ции solution аргумент start, чтобы подсчет начался с этого числа
	rand.Seed(time.Now().UnixNano())
	goal := 200
	cnt := 0 // счетчик сколько действий понадобилось
	for i := start; i != goal; {
		if i < goal {
			i = i + numbers[rand.Int31n(10)]
			fmt.Println(i)
			cnt++
		} else {
			i = i - numbers[rand.Int31n(10)]
			fmt.Println(i)
			cnt++
		}
	}
	return cnt
}
