package main

func main() {
	water := water{
		state: &LiquidState{
			State: &State{
				"液态",
			},
		},
		temperature: 20,
	}
	water.SetT(-4)
	water.Behavior()
	water.RaiseT(200)
	water.Behavior()
	water.ReduceT(300)
	water.Behavior()
}
