package main

type ObserverableM interface {
	AddObserver(ob Observer)
	RemoveObserver(ob Observer)
	NotifyObservers(tem interface{})
}

type Observerable struct {
	observer map[Observer]int
}

func (heater Observerable) AddObserver(ob Observer) {
	heater.observer[ob] = 1
}

func (heater Observerable) RemoveObserver(ob Observer) {
	delete(heater.observer, ob)
}

func (heater Observerable) NotifyObservers(tem interface{}) {
	for key, _ := range heater.observer {
		key.Update(tem)
	}
}
