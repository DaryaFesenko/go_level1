package calculator

import "errors"

type Calc struct {
	FirstNumber float32
	SecondNumber float32
	Result float32
	operations map[string] interface{}
}

func (c *Calc) Sum() (float32, error) {
	c.Result = c.FirstNumber + c.SecondNumber
	return c.Result, nil
}

func (c *Calc) Dif() (float32, error) {
	c.Result = c.FirstNumber - c.SecondNumber
	return c.Result, nil
}

func (c *Calc) Mult() (float32, error) {
	c.Result = c.FirstNumber * c.SecondNumber
	return c.Result, nil
}

func (c *Calc) Div() (float32, error) {
	if c.SecondNumber == 0{
		return 0, errors.New("divisor cannot be 0")
	}

	c.Result = c.FirstNumber / c.SecondNumber
	return c.Result, nil
}

func NewCalculator() *Calc {
	c := new(Calc)
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
		_, err := val.(func() (float32, error))()

		if err != nil{
			return 0, err
		}
	} else{
		return 0, errors.New("this operations not found")
	}

	return c.Result, nil
}
