package main

type deviceManager interface {
	Enumerate() deviceList
	Active(deviceId int)
	GetCurDeviceId() int
}

type speakerManager struct {
	curDeviceId int
}

func (s *speakerManager) Enumerate() deviceList {
	//TODO 真实项目应该通过驱动程序读取，此处模拟直接读取
	devices := deviceList{}
	devices.addDevice(deviceItem{
		id:          1,
		name:        "麦克风1",
		device_type: 1,
		isdefault:   false,
	})
	devices.addDevice(deviceItem{
		id:          2,
		name:        "麦克风2",
		device_type: 1,
		isdefault:   true,
	})
	return devices
}

func (s *speakerManager) Active(deviceId int) {
	s.curDeviceId = deviceId
}

func (s *speakerManager) GetCurDeviceId() int {
	return s.curDeviceId
}
