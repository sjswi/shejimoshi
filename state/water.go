package main

type water struct {
	*context
}

type wa interface {
	GetT() float64
	SetT(t float64)
	RaiseT(t float64)
	ReduceT(t float64)
	Behavior()
}

func (w *water) GetT() interface{} {
	return w.GetStateInfo()
}
func (w *water) SetT(t interface{}) {
	w.SetStateInfo(t)
}
func (w *water) RaiseT(t interface{}) {
	v := t.(float64)
	f := w.GetStateInfo().(float64)
	v = v + f
	w.SetT(v)

}
func (w *water) ReduceT(t interface{}) {
	v := t.(float64)
	f := w.GetStateInfo().(float64)
	v = f - v
	w.SetT(v)

}
func (w *water) Behavior() {
	w.GetState().Behavior()
}

//func (w *water) SetStateInfo(info interface{}) {
//
//
//	for key, _ := range w.states {
//		if key.IsMatch(info) {
//			err := w.ChangeState(key)
//			if err == nil {
//				return
//			}
//			fmt.Println(err)
//		}
//	}
//
//}
