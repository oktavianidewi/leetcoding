package main

import "fmt"

func convertTemperature(celsius float64) []float64 {
	res := make([]float64, 2)
	res[0] = celsius + 273.15
	res[1] = (celsius * 1.80) + 32.00
	return res
}

func main() {
	var celcius float64
	celcius = 36.50
	output := []float64{309.65000, 97.70000}
	res := convertTemperature(celcius)
	fmt.Printf("expected %v, result %v \n", output, res)
}
