package main

type deviceList struct {
	devices []deviceItem
}

func (receiver *deviceList) addDevice(device deviceItem) {
	if receiver.devices == nil {
		receiver.devices = make([]deviceItem, 1)
		receiver.devices[0] = device
		return
	}
	receiver.devices = append(receiver.devices, device)
}

func (receiver *deviceList) getCount() int {
	if receiver.devices == nil {
		return 0
	}
	return len(receiver.devices)
}

func (receiver *deviceList) getByIndex(index int) *deviceItem {
	if receiver.devices == nil || index < 0 || index >= len(receiver.devices) {
		return nil
	}
	return &receiver.devices[index]
}

func (receiver *deviceList) getById(id int) *deviceItem {
	for _, device := range receiver.devices {
		if device.getId() == id {
			return &device
		}
	}
	return nil
}
