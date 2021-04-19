package main

import (
	//calc "./calculator"
	"fmt"
	//fib "./fibonacci"
)

func main()  {

	a := 2
	b := &a
	c := &b
	fmt.Println(*b*b)
	/*
	// Первое и Второе задание с кэшом из мапы
	r := fib.Fibonacci(16)
	fmt.Println(r)
	r = fib.Fibonacci(4)
	fmt.Println(r)
	r = fib.Fibonacci(16)
	fmt.Println(r)
	r = fib.Fibonacci(4)
	fmt.Println(r)

	// Задания 3 и 4
	var num1, num2 float32
	var operation string

	fmt.Printf("Программа калькулятор\nПоддерживаются такие действия, как :\n   +    -    *   / \n")

	fmt.Print("Введите первое число : ")
	fmt.Scan(&num1)

	fmt.Print("Введите второе число : ")
	fmt.Scan(&num2)

	fmt.Print("Введите действие : ")
	fmt.Scan(&operation)

	res, err := calc.Calculator(num1, num2, operation)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}*/
}

