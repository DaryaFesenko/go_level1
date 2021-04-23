package calculatorFloat

import "errors"

type CalcFloat struct {
	FirstNumber float32
	SecondNumber float32
	Result float32
	operations map[string] interface{}
}

func (c CalcFloat) Sum() (interface{}, error) {
	c.Result = c.FirstNumber + c.SecondNumber
	return c.Result, nil
}

func (c CalcFloat) Dif() (interface{}, error) {
	c.Result = c.FirstNumber - c.SecondNumber
	return c.Result, nil
}

func (c CalcFloat) Mult() (interface{}, error) {
	c.Result = c.FirstNumber * c.SecondNumber
	return c.Result, nil
}

func (c CalcFloat) Div() (interface{}, error) {
	if c.SecondNumber == 0{
		return 0, errors.New("divisor cannot be 0")
	}

	c.Result = c.FirstNumber / c.SecondNumber
	return c.Result, nil
}

func NewCalculator() *CalcFloat {
	c := new(CalcFloat)
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

func Calculator(num1, num2 float32, operation string)  (float32, error){

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
