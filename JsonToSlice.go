package main

import (
	"encoding/json"
	"fmt"
)

type Friend struct {
	Name                  string `json:"name"`
	MobileNumber          string `json:"mobile_number"`
	AlternateMobileNumber string `json:"alternate_mobile_number,omitempty"`
	Address               string `json:"address"`
	City                  string `json:"city"`
}

func main() {
	friends := []Friend{
		{
			Name:                  "Alice",
			MobileNumber:          "123-456-7890",
			AlternateMobileNumber: "098-765-4321",
			Address:               "123 Main Street",
			City:                  "New York",
		},
		{
			Name:         "Bob",
			MobileNumber: "123-456-7891",
			Address:      "456 Market Street",
			City:         "Chicago",
		},
		{
			Name:                  "Jugaan",
			MobileNumber:          "123-456-7892",
			AlternateMobileNumber: "111-222-3333",
			Address:               "NY",
			City:                  "Los Angeles",
		},
		{
			Name:                  "Martha",
			MobileNumber:          "339-456-7892",
			AlternateMobileNumber: "568-222-3333",
			Address:               "789 North Mension",
			City:                  "Los Angeles",
		},
		{
			Name:                  "Charlie",
			MobileNumber:          "123-456-7892",
			AlternateMobileNumber: "111-222-3333",
			Address:               "789 Park Avenue",
			City:                  "Los Angeles",
		},
		{
			Name:         "Tony",
			MobileNumber: "143-456-4324",
			Address:      "Chicago",
			City:         "California",
		},
		{
			Name:                  "Alice",
			MobileNumber:          "123-456-7892",
			AlternateMobileNumber: "111-666-3333",
			Address:               "789 Park Avenue",
			City:                  "Los Angeles",
		},
		{
			Name:                  "Charlie",
			MobileNumber:          "123-456-7892",
			AlternateMobileNumber: "111-222-3333",
			Address:               "789 Park Avenue",
			City:                  "Los Angeles",
		}, {
			Name:         "Roman",
			MobileNumber: "123-556-6892",
			Address:      "789 Park Avenue",
			City:         "Los Angeles",
		},
	}

	jsonData, err := json.MarshalIndent(friends, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
	{

		var friends []Friend

		err := json.Unmarshal(jsonData, &friends)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(friends)

	}
}
