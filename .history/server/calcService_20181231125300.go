package server

import (
)

//CalcService calulator service
type CalcService struct{}

//AddValue add 2 values
func (s CalcService) AddValue(a,b fl) int {
	return a+b
}

//MultiplyValue multiply 2 values
func (s CalcService) MultiplyValue(a,b int) int {
	return a*b
}

func retDummyStr()string{
	return "dummy";
}