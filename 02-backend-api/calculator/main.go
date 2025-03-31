package calculator

import (
	"calculator/types"
)

func Add(num1 float64, num2 float64) *types.ResponseData {
	result := num1 + num2

	return &types.ResponseData{Result: result}
}

func Divide(num1 float64, num2 float64) *types.ResponseData {
	if num2 == 0 {
		return &types.ResponseData{Error: "Can not divide by 0"}
	}
	result := num1 / num2

	return &types.ResponseData{Result: result}
}

func Multiply(num1 float64, num2 float64) *types.ResponseData {
	result := num1 * num2

	return &types.ResponseData{Result: result}
}

func Subtract(num1 float64, num2 float64) *types.ResponseData {
	result := num1 - num2

	return &types.ResponseData{Result: result}
}

func Sum(items []float64) *types.ResponseData {
	var result float64 = 0
	for _, num := range items {
		result += num
	}

	return &types.ResponseData{Result: result}
}
