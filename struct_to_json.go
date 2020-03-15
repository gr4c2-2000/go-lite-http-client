package httpLiteClient

import (
	"encoding/json"
	"github.com/fatih/structs"
)

func InterfaceToJson(request interface{}) (jsonBytes []byte, err error) {
	requestMap := structs.Map(request)
	jsonBytes, err = json.Marshal(requestMap)
	if err != nil {
		return []byte{}, err
	}
	return jsonBytes, nil
}
