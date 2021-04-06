package main

func primeNumbers(to int) []int {
	var res []int

	for i := 2; i <= to; i++{
		if checkPrime(i) {
			res = append(res, i)
		}
	}

	return res
}


func checkPrime(num int) bool{
	for i := 2; i <= num; i++{
		if num % i == 0 && num != i{
			return false
		}
	}

	return true
}