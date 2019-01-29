package server

import (
)

type CalcService struct{}

func (s CalcService) AddValue(a,b int) int {
	return a+b
}

func (s CalcService) MultiplyValue(a,b int) int {
	return a*b
}

func retDummyStr()string{
	return "dummy";
}