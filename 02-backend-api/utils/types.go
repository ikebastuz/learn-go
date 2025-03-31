package utils

type RequestData2Nums struct {
	Number1 *float64 `json:"number1"`
	Number2 *float64 `json:"number2"`
}

type RequestDataNumsSlice struct {
	Items *[]float64 `json:"items"`
}

type ResponseData struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}
