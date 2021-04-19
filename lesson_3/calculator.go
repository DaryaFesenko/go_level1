package main

import "errors"

func calculator(num1, num2 float32, operation string) (float32, error){
	switch operation {
	case "+":
		return sum(num1, num2), nil
	case "-":
		return dif(num1, num2), nil
	case "*":
		return mult(num1, num2), nil
	case "/":
		return div(num1, num2)
	default:
		return 0, errors.New("the operation entered does not exist or is not supported")
	}
}

func sum(num1, num2 float32)  float32{
	return num1 + num2
}

func dif(num1, num2 float32)  float32{
	return num1 - num2
}

func mult(num1, num2 float32)  float32{
	return num1 * num2
}

func div(num1, num2 float32)  (float32, error){
	if num2 == 0{
		return 0, errors.New("divisor cannot be 0")
	}

	return num1 / num2, nil
}