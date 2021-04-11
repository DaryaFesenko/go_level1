package main

import (
	"fmt"
	"math/rand"
)

const arrayLen = 10


func main() {

	// Вариант с массивом
	array := [arrayLen]int{}

	for i := range array{
		array[i] = rand.Intn(100)
	}
	fmt.Println(array)
	array = sortArray(array)
	fmt.Println(array)

	// Вариант со слайсом
	slice := make([]int, 10)

	for i := range slice{
		slice[i] = rand.Intn(100)
	}

	fmt.Println(slice)
	sortSlice(slice)
	fmt.Println(slice)
}


// Если нужно как в задании отдавать на выходе,
// то лучше массив использовать, мне кажется
func sortArray(array [arrayLen]int) [arrayLen]int{
	for i := 1; i < len(array); i++{
		key := array[i]
		j := i - 1
		for j >= 0 && array[j] > key {
			array[j + 1] = array[j]
			j--
		}
		array[j+1] = key
	}

	return array
}

// Но ведь можно использовать слайс
func sortSlice(array []int) {
	for i := 1; i < len(array); i++{
		key := array[i]
		j := i - 1
		for j >= 0 && array[j] > key{
			array[j + 1] = array[j]
			j--
		}
		array[j+1] = key
	}
}
