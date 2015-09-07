package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func removeDuplicates(a []string) []string {
	result := []string{}
	seen := map[string]string{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}

	return result
}

func main() {

	jsonString := os.Args[1]

	var values []string
	json.Unmarshal([]byte(jsonString), &values)

	values = removeDuplicates(values)

	outputJSON, _ := json.Marshal(values)

	fmt.Println(string(outputJSON))
}
