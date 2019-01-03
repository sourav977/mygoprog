package main

import (
	"encoding/json" // Encoding and Decoding Package
	"fmt"
)

func main() {
	// Create a map of key/value pairs and parses the data into JSON
	emp := make(map[string]interface{})
	emp["name"] = "Sourav Patnaik"
	emp["age"] = 24
	//emp["1bc"] = 2
	emp["jobtitle"] = "Software Developer"
	emp["phone"] = map[string]interface{}{
		"home":   9090809099,
		"office": 1238906723,
	}
	emp["email"] = "sourav@gmail.com"

	// Marshal the map into a JSON string.
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonStr := string(empData)
	fmt.Println("The JSON data is:")
	fmt.Println(jsonStr)

	emp1 := make(map[string]interface{})
	json.Unmarshal(empData, &emp1)
	fmt.Println(emp1)
}
