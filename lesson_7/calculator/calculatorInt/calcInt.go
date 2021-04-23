package calculatorInt

import "errors"

type CalcInt struct {
	FirstNumber int
	SecondNumber int
	Result int
	operations map[string] interface{}
}

func (c CalcInt) Sum() (interface{}, error){
	c.Result = c.FirstNumber + c.SecondNumber
	return c.Result, nil
}

func (c CalcInt) Dif() (interface{}, error){
	c.Result = c.FirstNumber - c.SecondNumber
	return c.Result, nil
}

func (c CalcInt) Mult() (interface{}, error){
	c.Result = c.FirstNumber * c.SecondNumber
	return c.Result, nil
}

func (c CalcInt) Div() (interface{}, error){
	if c.SecondNumber == 0{
		return 0, errors.New("divisor cannot be 0")
	}

	c.Result = c.FirstNumber / c.SecondNumber
	return c.Result, nil
}

func NewCalculator() *CalcInt {
	c := new(CalcInt)
	c.operations = map[string] interface{} {
		"sum" : c.Sum,
		"+" : c.Sum,
		"dif": c.Dif,
		"-": c.Dif,
		"mult" : c.Mult,
		"*" : c.Mult,
		"div": c.Div,
		"/": c.Div,
	}

	return c
}

func Calculator(num1, num2 int, operation string)  (int, error){

	c := NewCalculator()
	c.FirstNumber = num1
	c.SecondNumber = num2

	if val, ok := c.operations[operation]; ok{
		err := val.(func() error)()

		if err != nil{
			return 0, err
		}
	}

	return c.Result, nil
}
