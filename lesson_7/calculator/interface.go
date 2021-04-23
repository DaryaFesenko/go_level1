package calculator

type Calculator interface {
	Sum() (interface{}, error)
	Dif() (interface{}, error)
	Div() (interface{}, error)
	Mult() (interface{}, error)
}