package handlers

import (
	"encoding/json"

	"net/http"

	. "../dataloaders"
	. "../models"
)

func Run() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	page := Page{ID: 7, Name: "Kullanıcılar", Description: "Kullanıcı Listesi", URI: "/users"}
	users := LoadUsers()
	interests := LoadInterests()
	interestMappings := LoadInterestMapings()
	var newusers []User
	for _, user := range users {
		for _, interestmapping := range interestMappings {
			if interestmapping.UserID == user.ID {
				for _, interest := range interests {
					if interest.ID == interestmapping.InterestID {
						user.Interests = append(user.Interests, interest)
					}
				}

			}
		}
		newusers = append(newusers, user)
	}
	ViewModel := UserViewModel{Page: page, Users: newusers}
	data, _ := json.Marshal(ViewModel)
	w.Write([]byte(data))
}
