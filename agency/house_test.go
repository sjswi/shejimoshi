package main

import "testing"

func TestName(t *testing.T) {
	myHouse := houseAgency{name: "我爱我家",
		houseInfos: make([]houseInfo, 10)}
	zhangsan := houseOwner{name: "张三",
		houseinfo: nil}
	lisi := houseOwner{name: "李四",
		houseinfo: nil}
	wangwu := houseOwner{name: "李四",
		houseinfo: nil}
	zhangsan.setHouseInfo(&houseInfo{
		area:        "20",
		price:       2500,
		hasWindow:   true,
		hasBathroom: true,
		hasKitchen:  false,
		address:     "城市家园",
		owner:       zhangsan,
	})
	lisi.setHouseInfo(&houseInfo{
		area:        "16",
		price:       1800,
		hasWindow:   true,
		hasBathroom: true,
		hasKitchen:  false,
		address:     "上地西里",
		owner:       zhangsan,
	})
	wangwu.setHouseInfo(&houseInfo{
		area:        "18",
		price:       3000,
		hasWindow:   true,
		hasBathroom: true,
		hasKitchen:  true,
		address:     "故宫",
		owner:       zhangsan,
	})
	zhangsan.publishHouseInfo(&myHouse)
	lisi.publishHouseInfo(&myHouse)
	wangwu.publishHouseInfo(&myHouse)
	tony := customer{
		name: "tony",
	}
	houseInfos := tony.findHouse("", myHouse)
	app := tony.seeHouse(houseInfos)
	tony.signContract(app, myHouse, 1)
}
