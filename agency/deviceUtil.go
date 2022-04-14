package main

import "shejimoshi/agency/constant"

type deviceUtil struct {
	magrs map[constant.DeviceType]deviceManager
}

func (d *deviceUtil) GetDeviceMagr(ty constant.DeviceType) deviceManager {
	return d.magrs[ty]
}

func (d *deviceUtil) GetDeviceList(ty constant.DeviceType) deviceList {
	return d.magrs[ty].Enumerate()
}

func (d *deviceUtil) Active(ty constant.DeviceType, devId int) {
	d.magrs[ty].Active(devId)
}

func (d *deviceUtil) GetCurDeviceId(ty constant.DeviceType) int {
	return d.magrs[ty].GetCurDeviceId()
}

type dev interface {
	GetDeviceMagr(ty constant.DeviceType) deviceManager
	GetDeviceList(ty constant.DeviceType) deviceList
	Active(ty constant.DeviceType, devId int)
	GetCurDeviceId(ty constant.DeviceType) int
}
