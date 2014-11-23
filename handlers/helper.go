package handlers

import (
	"encoding/json"
	// "fmt"
)

type ErrorMsg struct {
	Error string "json:'Error:'"
}

type SuccessMsg struct {
	Success string "json:'Success:'"
}

type jsonConvertible interface{}

func JsonString(obj jsonConvertible) (s string) {
	jsonObj, err := json.Marshal(obj)

	if err != nil {
		s = ""
	} else {
		s = string(jsonObj)
	}
	return
}
