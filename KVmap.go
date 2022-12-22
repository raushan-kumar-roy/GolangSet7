package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := `{
		"Thinly Sliced Peeled Apples": "6 Cups",
		"sugar": "3/4 cup",
		"flour": "2 tablespoons",
		"cinamon": "1/4 teaspoon",
		"nutmeg": "1/8 tablespoon",
		"lemon juice": "1 tablespoon"
	}`

	var m map[string]interface{}
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range m {
		fmt.Printf("%s: %v\n", k, v)
	}
}
