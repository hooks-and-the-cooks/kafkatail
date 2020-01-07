package translator

import (
	"encoding/json"
	"fmt"
)

func Translate(binData []byte) map[string]string {
	var jsonMap map[string]string

	err := json.Unmarshal(binData, &jsonMap)

	if err != nil {
		fmt.Println("Cannot unmarshal the json to proto-buf")
		return nil
	}
	return jsonMap
}
