package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func main()  {
	// Первое задание
	var a, b float32

	fmt.Println("Введите первую сторону прямоугольника :")
	fmt.Scanln(&a)

	fmt.Println("Введите вторую сторону прямоугольника :")
	fmt.Scanln(&b)

	s, err := getSquareRectangle(a, b)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Printf("Площадь прямоугольника со сторонами a = %0.2f, b = %0.2f равна %0.2f\n", a, b, s)

	// Второе задание
	var S float64

	fmt.Println("Введите площадь круга :")
	fmt.Scanln(&S)

	d, errD:= getCircleDiameter(S)
	c, errC := getCircumference(S)

	if errD != nil{
		fmt.Println(errD)
	}

	if errC != nil{
		fmt.Println(errD)
	}

	fmt.Printf("Диаметр окружности равен %0.2f, длина окружности равна %0.2f\n", d, c)

	// Третье задание
	var number int

	fmt.Println("Введите трехзначное число : ")
	fmt.Scanln(&number)

	var digits, err3 = getDigitsOfNumber(number)

	if err3 != nil{
		fmt.Println(err3)
	}

	fmt.Printf("Сотни %d  Десятки %d  Единицы %d", digits[2], digits[1], digits[0])
}

// Функция для расчета площади прямоугольника по заданным сторонам
func getSquareRectangle(a, b float32) (float32, error) {
	if a < 0 || b < 0{
		return 0, errors.New("sides cannot be less than zero")
	}

	return a * b, nil
}

// Функция для расчета диаметра окружности по заданной площади круга
func getCircleDiameter(S float64) (float64, error) {
	if S < 0{
		return 0, errors.New("square cannot be less than zero")
	}

	return 2 * math.Sqrt(S / math.Pi), nil
}

// Функция расчета длины окружности по заданной площади круга
func getCircumference(S float64) (float64, error)  {
	if S < 0{
		return 0, errors.New("square cannot be less than zero")
	}
	d, err := getCircleDiameter(S)
	if err != nil{
		return 0, err
	}
	return d * math.Pi, nil
}

// Функция для разделения трехзначного числа на разряды
func getDigitsOfNumber(number int) ([3]int, error) {
	var digits [3]int

	if getCountsOfDigits(number) != 3{
		return digits, errors.New("the entered number is not three-digit")
	}

	for i := 0; i <= 2; i++{
		digits[i] = number % 10

		number = number / 10
	}

	return digits, nil
}

// Функция получения количества цифр в числе
func getCountsOfDigits(number int) int{
	return len(strconv.Itoa(number))
}