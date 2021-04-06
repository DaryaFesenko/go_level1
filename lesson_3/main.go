package main

import "fmt"

func main(){
	// Первое задание

	var num1, num2 float32
	var operation string

	fmt.Printf("Программа калькулятор\nПоддерживаются такие действия, как :\n   +    -    *   / \n")

	fmt.Print("Введите первое число : ")
	fmt.Scan(&num1)

	fmt.Print("Введите второе число : ")
	fmt.Scan(&num2)

	fmt.Print("Введите действие : ")
	fmt.Scan(&operation)

	res, err := calculator(num1, num2, operation)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}


    // Второе задание
    var num int

	fmt.Printf("Поиск простых чисел от 1 до N\n Введите N")
	fmt.Scan(&num)

	fmt.Print(primeNumbers(num))


}
