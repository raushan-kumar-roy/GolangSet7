package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	stateCapital := map[string]string{
		"Andhra Pradesh":    "Amaravati",
		"Arunachal Pradesh": "Itanagar",
		"Assam":             "Dispur",
		"Bihar":             "Patna",
		"Chhattisgarh":      "Raipur",
		"Goa":               "Panaji",
		"Gujarat":           "Gandhinagar",
		"Haryana":           "Chandigarh",
		"Himachal Pradesh":  "Shimla",
		"Jammu and Kashmir": "Srinagar",
	}
	jsonData, err := json.MarshalIndent(stateCapital, "", " ")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(jsonData)
	}

}
