package main

import "fmt"

type houseOwner struct {
	name      string
	houseinfo *houseInfo
}

func (house *houseOwner) getName() string {
	return house.name
}

func (house *houseOwner) setHouseInfo(info *houseInfo) {
	house.houseinfo = info
}

func (house *houseOwner) publishHouseInfo(agency *houseAgency) {
	agency.addHouseInfo(*house.houseinfo)
	fmt.Printf("%s发布了房源\n", house.getName())
}
