package main

func main() {
	heater := WaterHeater{
		ob:          Observerable{observer: make(map[Observer]int)},
		temperature: 0.0,
	}
	w := washingMode{}
	d := drinkingMode{}
	heater.addObserver(&w)
	heater.addObserver(&d)
	heater.setTemperature(40.0)
	heater.setTemperature(60.0)
	heater.setTemperature(100.0)
}
