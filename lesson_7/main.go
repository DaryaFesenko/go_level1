package main

import (
	"fmt"
	"reflect"
	"errors"

	"./calculator/calculatorInt"
	"./calculator/calculatorFloat"
	"./calculator/calculatorUint"

	"./calculator"
)


func main()  {
	var calcInt = calculatorInt.NewCalculator()
	var calcFloat = calculatorFloat.NewCalculator()
	var calcUint = calculatorUint.NewCalculator()

	CheckImplementedInterface(calcInt, calcFloat, calcUint)

	var (
		operatorInt string
		operatorFloat string
		operatorUint string
	)

	fmt.Scanln(&calcInt.FirstNumber, &operatorInt, &calcInt.SecondNumber)
	fmt.Scanln(&calcFloat.FirstNumber, &operatorFloat, &calcFloat.SecondNumber)
	fmt.Scanln(&calcUint.FirstNumber, &operatorUint, &calcUint.SecondNumber)

	var resultInt, errInt = Calculate(calcInt, operatorInt)
	if errInt != nil {
		fmt.Println(errInt)
	}
	fmt.Println(resultInt)


	var resultFloat, errFloat = Calculate(calcFloat, operatorFloat)
	if errFloat != nil {
		fmt.Println(errFloat)
	}
	fmt.Println(resultFloat)


	var resultUint, errUint = Calculate(calcUint, operatorUint)
	if errUint != nil {
		fmt.Println(errUint)
	}
	fmt.Println(resultUint)
}

func isImplemented(c interface{}) bool{
	if _, ok := c.(calculator.Calculator); ok{
		return true
	}
	return false
}

func CheckImplementedInterface(calc ...interface{}){
	for _, v := range calc{
		if !isImplemented(v) {
			fmt.Println(reflect.TypeOf(v).String() + " is not implemented Calculator")
			return
		}
	}
}

func Calculate(c calculator.Calculator, operator string) (interface{}, error){
	switch operator {
	case "+":
		return Sum(c)
		break
	case "-":
		return Dif(c)
		break
	case "*":
		return Mult(c)
		break
	case "/":
		return Div(c)
		break
	}

	return 0, errors.New("\nsuch operation does not exist or it has not yet been implemented")
}

func Sum(c calculator.Calculator) (interface{}, error){
	var res, err = c.Sum()
	return res, err
}

func Dif(c calculator.Calculator) (interface{}, error){
	var res, err = c.Dif()
	return res, err
}

func Div(c calculator.Calculator) (interface{}, error){
	var res, err = c.Div()
	return res, err
}

func Mult(c calculator.Calculator) (interface{}, error){
	var res, err = c.Mult()
	return res, err
}