package models

import "encoding/json"

type GAEBool int

func (b *GAEBool) MarshalJSON() ([]byte, error) {
	if *b == 1 {
		return json.Marshal(true)
	} else if *b == -1 {
		return json.Marshal(false)
	}
	return json.Marshal(0)
}

func (t *GAEBool) UnmarshalJSON(data []byte) error {
	strBool := string(data)
	if strBool == "true" {
		*t = 1
	} else if strBool == "false" {
		*t = -1
	}
	return nil
}

func (b *GAEBool) Bool() bool {
	if *b == 1 {
		return true
	}
	return false
}
