package dataloaders

import (
	"encoding/json"

	. "app/models"
	"app/utils"
)

func LoadUsers() []User {
	bytes, _ := utils.ReadFile("../json/users.json")
	var data []User
	json.Unmarshal([]byte(bytes), &data)
	return data
}
func LoadInterests() []Interest {
	bytes, _ := utils.ReadFile("../json/Interests.json")
	var data []Interest
	json.Unmarshal([]byte(bytes), &data)
	return data
}
func LoadInterestMapings() []InterestMapping {
	bytes, _ := utils.ReadFile("../json/InterestMappings.json")
	var data []InterestMapping
	json.Unmarshal([]byte(bytes), &data)
	return data
}
