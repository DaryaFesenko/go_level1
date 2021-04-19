package fibonacci

import "fmt"

type fibonacci struct {
	Cached map[int]int
	Number int
	Result int
}


func NewFibonacci()  *fibonacci{
	f := new(fibonacci)

	f.Cached = map[int]int{
		1 : 1,
		2 : 1,
	}

	return f
}

var f *fibonacci = NewFibonacci()

func Fibonacci(n int)  int{
	f.Number = n

	if val, ok := f.Cached[f.Number]; ok{
		defer func(val,n int) {
			fmt.Printf("Из кэша использовано: [%v] = %v\n", n, val)
		}(val, f.Number)
		return val
	}


	f.Result = calculationFibonacci(f.Number)
	f.Cached[f.Number] = f.Result
	return f.Result
}


func calculationFibonacci(n int) int{
	if n == 1 || n == 2 {
		return 1
	}

	return calculationFibonacci(n - 1) + calculationFibonacci(n - 2)
}
