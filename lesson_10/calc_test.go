package calculator

import (
	"errors"
	"testing"
	"fmt"
)

type In struct{
	First float32
	Second float32
	Operation string
}

func TestCalc(t *testing.T){
	testCases :=[]struct{
		Name string
		In In
		Out float32
	}{
		{ Name: "Sum", In : In{First: 5.1, Second: 4.6, Operation: "+"}, Out: 9.7},
		{ Name: "Div", In : In{First: -7.3, Second: 9.3, Operation: "-"}, Out: -16.6},
		{ Name: "Dif", In : In{First: 6.6, Second: -1.1, Operation: "/"}, Out: -6},
		{ Name: "Mult", In : In{First: 6.1, Second: -0.5, Operation: "*"}, Out: -3.05},
	}

	for _, tt := range testCases{
		t.Run(tt.Name, func(t *testing.T) {
			if got, _ := Calculator(tt.In.First, tt.In.Second, tt.In.Operation); got != tt.Out{
				t.Fatalf("got %f, but want %f", got, tt.Out)
			}
		})
	}
}

var result float32

func BenchmarkCalc(b *testing.B)  {
	var res float32

	for i:=0; i< b.N; i++{
		res, _ = Calculator(4.55, 6.777, "*")
	}

	result = res
}

func TestFuzzyCalc(t *testing.T){
	testDif0 :=struct{
		Name string
		In In
		Out error
	}{ Name: "Dif 0", In : In{First: 5.1, Second: 0, Operation: "/"}, Out: errors.New("divisor cannot be 0")}

	testOperNotFound :=struct{
		Name string
		In In
		Out error
	}{ Name: "Operation", In : In{First: -7.3, Second: 9.3, Operation: "87"}, Out: errors.New("this operations not found")}

	t.Run(testDif0.Name, func(t *testing.T){
		if _, err := Calculator(testDif0.In.First, testDif0.In.Second, testDif0.In.Operation);  err.Error() != testDif0.Out.Error() {
			t.Fatalf("an error was expected: %s", testDif0.Out)
		}
	})

	t.Run(testOperNotFound.Name, func(t *testing.T){
		if _, err := Calculator(testOperNotFound.In.First, testOperNotFound.In.Second, testOperNotFound.In.Operation); err.Error() != testOperNotFound.Out.Error() {
			fmt.Println(err)
			t.Fatalf("an error was expected: %s", testOperNotFound.Out)
		}
	})

}