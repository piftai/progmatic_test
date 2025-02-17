package service

import (
	"math/rand"
	"time"
)

var numbers = [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

//  в процессе решения пришел к тому, что можно сделать более гибкое решение, и передавать в solution
// помимо начального значения еще и конечное, чтобы можно было посчитать любое число в принципе,
// но этого не указано в условиях

func Solution(start int) int { // сделаем у ф-ции solution аргумент start, чтобы подсчет начался с этого числа
	rand.Seed(time.Now().UnixNano()) // с 1.24 не работает
	goal := 200
	cnt := 0 // счетчик сколько действий понадобилось
	for i := start; i != goal; {
		if i < goal {
			i = i + numbers[rand.Int31n(10)]
			// fmt.Println(i) you can uncomment it to watch the progress
			cnt++
		} else {
			i = i - numbers[rand.Int31n(10)]
			// fmt.Println(i) you can uncomment it to watch the progress
			cnt++
		}
	}
	return cnt
}
