package main

import "fmt"

type houseInfo struct {
	area        string
	price       float64
	hasWindow   bool
	hasBathroom bool
	hasKitchen  bool
	address     string
	owner       houseOwner
}

func (house *houseInfo) getAddress() string {
	return house.address
}
func (house *houseInfo) getOwnerName() string {
	return house.owner.getName()
}
func (house *houseInfo) showInfo() {

	window := "无"
	bath := "无"
	kit := "无"

	if house.hasBathroom {
		bath = "有"
	}
	if house.hasKitchen {
		kit = "有"
	}
	if house.hasWindow {
		window = "有"
	}
	fmt.Printf("面积：%s\n价格：%.2f\n窗户：%s\n卫生间：%s\n厨房：%s\n地址：%s\n房东：%s\n",
		house.area, house.price, window, bath, kit, house.address, house.getOwnerName())
}
