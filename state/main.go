package main

import "sync"

var once sync.Once
var instance1 *SolidState

func GetInstance1() *SolidState {
	once.Do(func() {
		instance1 = new(SolidState)
	})
	return instance1
}

var instance2 *LiquidState

func GetInstance2() *LiquidState {
	once.Do(func() {
		instance2 = new(LiquidState)
	})
	return instance2
}

var instance3 *GaseousState

func GetInstance3() *GaseousState {
	once.Do(func() {
		instance3 = new(GaseousState)
	})
	return instance3
}
func main() {

	water := water{
		context: &context{
			states: map[st]int{
				GetInstance1(): 1,
				GetInstance2(): 1,
				GetInstance3(): 1,
			},
			cur_state:  nil,
			state_info: float64(0),
		},
	}
	water.SetT(23.0)
	water.Behavior()
	water.RaiseT(200.0)
	water.Behavior()
	water.ReduceT(300.0)
	water.Behavior()
}
