package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func PrettyRequest(req string) string {
	s := req
	s = strings.ReplaceAll(s, " ", "")
	return s
}

func ParseMap[T any](m map[string]any, v T) error {
	b, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("error marshaling map: %s", err.Error())
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return fmt.Errorf("error unmarshaling map: %s", err.Error())
	}

	return nil
}

func SaveToFile(name string, data []byte) error {
	err := os.WriteFile(name, data, os.ModePerm)
	return err
}

func LoadQuery(file string) string {
	queryBytes, err := os.ReadFile(file)
	if err != nil {
		panic(fmt.Sprintf("Error reading query file. File: %s, Error: %s", file, err.Error()))
	}
	return string(queryBytes)
}

func LoadParams(file string) map[string]any {
	inputBytes, err := os.ReadFile(file)
	if err != nil {
		panic(fmt.Sprintf("Error reading .json file: File: %s, Error: %s", file, err.Error()))
	}
	input := map[string]any{}
	err = json.Unmarshal(inputBytes, &input)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshaling .json file: Error: %s", err.Error()))
	}

	return input
}
