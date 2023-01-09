package dataloaders

import (
	"encoding/json"

	. "../models"
	util "../utils"
)

func LoadUsers() []User {
	bytes, _ := util.ReadFile("../json/users.json")
	var data []User
	json.Unmarshal([]byte(bytes), &data)
	return data
}
func LoadInterests() []Interest {
	bytes, _ := util.ReadFile("../json/Interests.json")
	var data []Interest
	json.Unmarshal([]byte(bytes), &data)
	return data
}
func LoadInterestMapings() []InterestMapping {
	bytes, _ := util.ReadFile("../json/InterestMappings.json")
	var data []InterestMapping
	json.Unmarshal([]byte(bytes), &data)
	return data
}
