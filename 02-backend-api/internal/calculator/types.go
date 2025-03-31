package calculator

import (
	"fmt"
)

// Operation represents a mathematical operation
type Operation string

const (
	OpAdd      Operation = "add"
	OpSubtract Operation = "subtract"
	OpMultiply Operation = "multiply"
	OpDivide   Operation = "divide"
	OpSum      Operation = "sum"
)

// TwoNumberRequest represents a request with two numbers
type TwoNumberRequest struct {
	Number1 int `json:"number1" validate:"required"`
	Number2 int `json:"number2" validate:"required"`
}

// SumRequest represents a request with an array of numbers
type SumRequest struct {
	Numbers []int `json:"items" validate:"required,min=1"`
}

// Response represents the API response
type Response struct {
	Result int    `json:"result"`
	Error  string `json:"error,omitempty"`
}

// Validate performs validation on the request
func (r *TwoNumberRequest) Validate() error {
	if r.Number1 == 0 && r.Number2 == 0 {
		return fmt.Errorf("both numbers cannot be zero")
	}
	return nil
}

// Validate performs validation on the sum request
func (r *SumRequest) Validate() error {
	if len(r.Numbers) == 0 {
		return fmt.Errorf("numbers array cannot be empty")
	}
	return nil
} 