package tools

import (
	"fmt"
	"encoding/json"
)

func ReturnJsonIndent(jsonstruct interface{}) (string, error){

	bytes, err := json.MarshalIndent(jsonstruct, "", "    ")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}

	return string(bytes), nil
}

func ReturnJsonString(jsonstruct interface{}) (string, error){

	bytes, err := json.Marshal(jsonstruct)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}

	return string(bytes), nil
}