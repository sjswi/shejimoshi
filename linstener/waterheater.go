package main

type WaterHeater struct {
	ob          Observerable
	temperature float64
}

func (heater *WaterHeater) setTemperature(temperature float64) {
	heater.temperature = temperature
	heater.notify()
}

func (heater *WaterHeater) getTemperature() float64 {
	return heater.temperature
}

func (heater *WaterHeater) addObserver(obser Observer) {
	heater.ob.AddObserver(obser)
}
func (heater *WaterHeater) removeObserver(obser Observer) {
	heater.ob.RemoveObserver(obser)
}
func (heater *WaterHeater) notify() {
	heater.ob.NotifyObservers(heater.temperature)
}
