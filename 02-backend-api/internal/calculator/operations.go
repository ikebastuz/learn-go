package calculator

import (
	"fmt"
)

// Add performs addition of two numbers
func Add(req TwoNumberRequest) (Response, error) {
	if err := req.Validate(); err != nil {
		return Response{Error: err.Error()}, err
	}
	return Response{Result: req.Number1 + req.Number2}, nil
}

// Subtract performs subtraction of two numbers
func Subtract(req TwoNumberRequest) (Response, error) {
	if err := req.Validate(); err != nil {
		return Response{Error: err.Error()}, err
	}
	return Response{Result: req.Number1 - req.Number2}, nil
}

// Multiply performs multiplication of two numbers
func Multiply(req TwoNumberRequest) (Response, error) {
	if err := req.Validate(); err != nil {
		return Response{Error: err.Error()}, err
	}
	return Response{Result: req.Number1 * req.Number2}, nil
}

// Divide performs division of two numbers
func Divide(req TwoNumberRequest) (Response, error) {
	if err := req.Validate(); err != nil {
		return Response{Error: err.Error()}, err
	}
	if req.Number2 == 0 {
		return Response{Error: "division by zero is not allowed"}, fmt.Errorf("division by zero")
	}
	return Response{Result: req.Number1 / req.Number2}, nil
}

// Sum adds all numbers in the array
func Sum(req SumRequest) (Response, error) {
	if err := req.Validate(); err != nil {
		return Response{Error: err.Error()}, err
	}
	result := 0
	for _, num := range req.Numbers {
		result += num
	}
	return Response{Result: result}, nil
} 