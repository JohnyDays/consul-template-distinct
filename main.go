package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func removeDuplicates(a map[string]string) []string {
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

	var values map[string]string
	json.Unmarshal([]byte(jsonString), &values)

	distinctValues := removeDuplicates(values)

	outputJSON, _ := json.Marshal(distinctValues)

	fmt.Println(string(outputJSON))
}
