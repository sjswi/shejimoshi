package main

import (
	"fmt"
	"shejimoshi/agency/constant"
	"testing"
)

func Test(t *testing.T) {
	deviceutil := deviceUtil{
		magrs: map[constant.DeviceType]deviceManager{
			constant.TypeSpeaker: &speakerManager{},
		},
	}
	devi := deviceutil.GetDeviceList(constant.TypeSpeaker)
	fmt.Printf("麦克风设备\n")
	if devi.getCount() > 0 {
		deviceutil.Active(constant.TypeSpeaker, devi.getByIndex(1).getId())
	}
	for i := 0; i < devi.getCount(); i++ {
		fmt.Printf("%s\n", devi.getByIndex(i).toString())
	}
	fmt.Printf("当前使用的设备:%s\n", devi.getById(deviceutil.GetCurDeviceId(constant.TypeSpeaker)).getName())
}
