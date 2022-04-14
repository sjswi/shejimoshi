package main

import "fmt"

type customer struct {
	name string
}

func (cu customer) getName() string {
	return cu.name
}
func (cu customer) findHouse(desc string, agency houseAgency) []houseInfo {
	fmt.Printf("%s想找一个%s的房子\n", cu.getName(), desc)
	return agency.getMatchInfo(desc)
}
func (cu customer) seeHouse(house []houseInfo) houseInfo {

	//TODO 用户选择
	size := len(house)
	return house[size-1]
}
func (cu customer) signContract(house houseInfo, agency houseAgency, number int) {
	//签约
	fmt.Printf("%s与中介%s签约%d个月\n", cu.getName(), agency.getName(), number)
	agency.signContract(house)
}
