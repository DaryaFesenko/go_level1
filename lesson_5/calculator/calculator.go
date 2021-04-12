package calculator

import "errors"

type calculator struct {
	FirstNumber float32
	SecondNumber float32
	Result float32
}

func (c *calculator) Sum() error{
	c.Result = c.FirstNumber + c.SecondNumber
	return nil
}

func (c *calculator) Dif() error{
	c.Result = c.FirstNumber - c.SecondNumber
	return nil
}

func (c *calculator) Mult() error{
	c.Result = c.FirstNumber * c.SecondNumber
	return nil
}

func (c *calculator) Div() error{
	if c.SecondNumber == 0{
		return errors.New("divisor cannot be 0")
	}

	c.Result = c.FirstNumber / c.SecondNumber
	return nil
}


var operations map[string] interface{}
var c calculator

func init(){
	operations = map[string] interface{} {
		"sum" : c.Sum,
		"+" : c.Sum,
		"dif": c.Dif,
		"-": c.Dif,
		"mult" : c.Mult,
		"*" : c.Mult,
		"div": c.Div,
		"/": c.Div,
	}
}

// Возможно, стоило все таки упростить и просто переопределить методы Sum, Dif и тд.
// Но тогда я не до конца пойму, зачем структуры, когда можно было 4 простых функции в пакет положить
// И зачем вы показывали мапу с функциями тогда


func Calculator(num1, num2 float32, operation string)  (float32, error){
	c.FirstNumber = num1
	c.SecondNumber = num2

	if val, ok := operations[operation]; ok{
		err := val.(func() error)()

		if err != nil{
			return 0, err
		}
	}

	return c.Result, nil
}
