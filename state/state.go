package main

import "fmt"

type st interface {
	GetName() string
	Behavior()
	IsMatch(info interface{}) bool
}
type State struct {
	name string
}

func (s *State) GetName() string {
	return s.name
}
func (s *State) Behavior() {
	//需要子类重写
}
func (s *State) IsMatch(info interface{}) bool {
	//是否在当前状态的范围内
	//同样需要子类重写
	return false
}

type SolidState struct {
	*State
}

func (s *SolidState) IsMatch(info interface{}) bool {
	//是否在当前状态的范围内
	//同样需要子类重写
	f := info.(float64)
	if f < 0 {
		return true
	} else {
		return false
	}
}
func (s *SolidState) Behavior() {
	fmt.Println("固态")
}

type GaseousState struct {
	*State
}

func (s *GaseousState) IsMatch(info interface{}) bool {
	//是否在当前状态的范围内
	//同样需要子类重写
	f := info.(float64)
	if f > 100 {
		return true
	} else {
		return false
	}
}
func (s *GaseousState) Behavior() {
	fmt.Println("气态")
}
func (s *LiquidState) IsMatch(info interface{}) bool {
	//是否在当前状态的范围内
	//同样需要子类重写
	f := info.(float64)
	if f >= 0 && f <= 100 {
		return true
	} else {
		return false
	}
}

type LiquidState struct {
	*State
}

func (s *LiquidState) Behavior() {
	fmt.Println("液态")
}
