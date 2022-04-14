package main

import (
	"shejimoshi/agency/constant"
	"strconv"
)

type deviceItem struct {
	id          int
	name        string
	device_type constant.DeviceType
	isdefault   bool
}

func (receiver *deviceItem) toString() string {
	s := "type:" + string(rune(receiver.device_type)) + "id:" + string(rune(receiver.id)) + "name:" + receiver.name + "isDefault" + strconv.FormatBool(receiver.isdefault)
	return s
}

func (receiver *deviceItem) getId() int {
	return receiver.id
}

func (receiver *deviceItem) getName() string {
	return receiver.name
}

func (receiver *deviceItem) getType() constant.DeviceType {
	return receiver.device_type
}
func (receiver *deviceItem) isDefault() bool {
	return receiver.isdefault
}
