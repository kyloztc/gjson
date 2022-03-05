package gjson

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ExtractValueFromJson 从json字符串中取出指定key值
func ExtractValueFromJson(jsonStr string, key string) (interface{}, error) {
	var jsonInterface interface{}
	err := json.Unmarshal([]byte(jsonStr), &jsonInterface)
	if err != nil {
		return nil, error_json_format_error
	}
	keyField := strings.Split(key, ".")
	return getValueFromJsonInterface(jsonInterface, keyField)
}

func getValueFromJsonInterface(jsonInterface interface{}, fieldKey []string) (interface{}, error) {
	key := fieldKey[0]
	var processRsp interface{}
	var err error
	switch ji := jsonInterface.(type) {
	case []interface{}:
		processRsp, err = parseListFromInterface(ji, key)
	default:
		processRsp, err = parseMapFromInterface(ji, key)
	}
	if len(fieldKey) == 1 || err != nil {
		return processRsp, err
	}
	return getValueFromJsonInterface(processRsp, fieldKey[1:])
}

func parseListFromInterface(listInterface []interface{}, key string) (interface{}, error) {
	rsp := make([]interface{}, 0)
	for _, list := range listInterface {
		mapInfo, ok := list.(map[string]interface{})
		if !ok {
			return nil, errors.New(fmt.Sprintf("type error, %s is not array", key))
		}
		distValue, ok := mapInfo[key]
		if !ok {
			continue
		}
		rsp = append(rsp, distValue)
	}
	return rsp, nil
}

func parseMapFromInterface(mapInterface interface{}, key string) (interface{}, error) {
	mapInfo, ok := mapInterface.(map[string]interface{})
	if !ok {
		return nil, errors.New(fmt.Sprintf("type error, %s is not map", key))
	}
	rsp, ok := mapInfo[key]
	if !ok {
		return nil, nil
	}
	return rsp, nil
}
