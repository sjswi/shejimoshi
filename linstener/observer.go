package main

import (
	"fmt"
)

type Observer interface {
	Update(tem interface{})
}

type washingMode struct {
}
type drinkingMode struct {
}

func (a *washingMode) Update(tem interface{}) {

	f := tem.(float64)
	if f >= 50 && f <= 70 {
		fmt.Println("水温正好，可以洗澡")
	}
}

func (b *drinkingMode) Update(tem interface{}) {

	f := tem.(float64)
	if f == 100 {
		fmt.Println("可以喝了")
	}
}
