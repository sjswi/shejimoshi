package main

import "fmt"

type st interface {
	GetName() string
	Behavior()
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

type SolidState struct {
	*State
}

func (s *SolidState) Behavior() {
	fmt.Println("固态")
}

type GaseousState struct {
	*State
}

func (s *GaseousState) Behavior() {
	fmt.Println("气态")
}

type LiquidState struct {
	*State
}

func (s *LiquidState) Behavior() {
	fmt.Println("液态")
}

type water struct {
	temperature float64
	state       st
}

type wa interface {
	SetState(state st)
	ChangeState(state st)
	GetT() float64
	SetT(t float64)
	RaiseT(t float64)
	ReduceT(t float64)
	Behavior()
}

func (w *water) SetState(state *st) {
	w.state = *state
}
func (w *water) ChangeState(state st) {
	w.state = state
}

func (w *water) GetT() float64 {
	return w.temperature
}
func (w *water) SetT(t float64) {
	w.temperature = t
	if t < 0 {

		w.ChangeState(&SolidState{State: &State{
			name: "固态",
		}})
	} else if t < 100 {
		w.ChangeState(&LiquidState{State: &State{
			name: "液态",
		}})
	} else {
		w.ChangeState(&GaseousState{State: &State{
			name: "气态",
		}})
	}
}
func (w *water) RaiseT(t float64) {
	w.SetT(w.temperature + t)

}
func (w *water) ReduceT(t float64) {
	w.SetT(w.temperature - t)

}
func (w *water) Behavior() {
	w.state.Behavior()
}
