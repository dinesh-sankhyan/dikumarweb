package server

import (
)

//CalcService calulator service
type CalcService struct{}

//AddValue add 2 values
func (s CalcService) AddValue(a,b int) int {
	return a+b
}

////AddValue add 2 values
func (s CalcService) MultiplyValue(a,b int) int {
	return a*b
}

func retDummyStr()string{
	return "dummy";
}