package main

import (
	"encoding/json"
	"fmt"

	"github.com/bgadrian/fastfaker/faker"
)

/**
Generate random custom user data to mock/test/stress your Auth/User systems.
*/
func main() {
	count := 12

	//for all the variables available see:
	//https://github.com/bgadrian/fastfaker/TEMPLATE_VARIABLES.md
	type User struct {
		ID            string `json:"id" fake:"{uuid}"`
		Username      string `fake:"{username}"`
		Password      string `fake:"{passwordfull}"`
		Avatar        string `fake:"{avatarurl}"`
		Name          string `fake:"{name}"`
		Address       string `fake:"{street} {city} {country}"`
		FavoriteColor string `fake:"{safecolor}"`
		Bio           string `fake:"{sentenceavg}"`
		Twitter       string `fake:"{word}##"`
		Website       string `fake:"{url}"`
		RegisterData  string `fake:"{datecurrentyearstr}"`
	}

	for i := 0; i < count; i++ {
		newUser := &User{}
		faker.Global.Struct(newUser)

		asJSON, _ := json.Marshal(newUser)
		fmt.Println(string(asJSON))
	}
}
