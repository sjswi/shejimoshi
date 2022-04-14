package main

import "fmt"

type houseAgency struct {
	houseInfos []houseInfo
	name       string
}

func (house *houseAgency) getName() string {
	return house.name
}

func (house *houseAgency) addHouseInfo(houseinfo houseInfo) {
	if house.houseInfos == nil {
		house.houseInfos = make([]houseInfo, 1)
		house.houseInfos[0] = houseinfo
		return
	}
	house.houseInfos = append(house.houseInfos, houseinfo)

}

func (house *houseAgency) removeHouseInfo(houseinfo houseInfo) {
	for i, info := range house.houseInfos {
		if info == houseinfo {
			house.houseInfos = append(house.houseInfos[:i], house.houseInfos[i+1:]...)
			break
		}
	}

}

func (house *houseAgency) getSearchCondition(description string) string {
	//TODO 将描述转换为可搜索的条件,
	return description
}

func (house *houseAgency) getMatchInfo(desc string) []houseInfo {
	//TODO 从可搜索的条件搜索房源信息并返回
	for _, info := range house.houseInfos {
		info.showInfo()
	}
	return house.houseInfos
}

func (house *houseAgency) signContract(houseinfo houseInfo) {
	fmt.Printf("%s与房东%s签定合约\n", house.getName(), houseinfo.getOwnerName())
	house.removeHouseInfo(houseinfo)
}

func (house *houseAgency) signContracts() {
	for _, houseinfo := range house.houseInfos {
		house.signContract(houseinfo)
	}
}
