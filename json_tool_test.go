package gjson

import (
	"fmt"
	"testing"
)

func TestExtractValueFromJson(t *testing.T) {
	// 列表类型字符串
	listJsonStr := "[{\"key0\": 1, \"key1\": \"v1\"}, {\"key0\": 2, \"key1\": \"v2\"}, {\"key0\": 3, \"key1\": \"v1\"}]"
	listDistKey := "key0"
	listDistValue, err := ExtractValueFromJson(listJsonStr, listDistKey)
	if err != nil {
		t.Errorf("test list json error|%v\n", err)
		return
	}
	fmt.Printf("list dist value: %v\n", listDistValue)

	// map/struct类型json字符串
	mapJsonStr := "{\"key0\": 1, \"key1\": \"v1\", \"key2\": 1.1}"
	mapDistKey := "key2"
	mapJsonDistValue, err := ExtractValueFromJson(mapJsonStr, mapDistKey)
	if err != nil {
		t.Errorf("test map json error|%v", err)
		return
	}
	fmt.Printf("map dist value: %v\n", mapJsonDistValue)

	// 嵌套key类型
	nestJsonStr := "{\"root\": {\"sub\": {\"key\": \"distValue\"}}}"
	nestDistKey := "root.sub.key"
	nestDistValue, err := ExtractValueFromJson(nestJsonStr, nestDistKey)
	if err != nil {
		t.Errorf("test nest json error|%v", err)
		return
	}
	fmt.Printf("nest dist value: %v\n", nestDistValue)

	nestListJsonStr := "[{\"root\": {\"sub\": {\"key\": \"distValue0\"}}}, " +
		"{\"root\": {\"sub\": {\"key\": \"distValue1\"}}}, " +
		"{\"root\": {\"sub\": {\"key\": \"distValue2\"}}}]"
	nestListDistValue, err := ExtractValueFromJson(nestListJsonStr, nestDistKey)
	if err != nil {
		t.Errorf("test nest json error|%v", err)
		return
	}
	fmt.Printf("nest list dist value: %v\n", nestListDistValue)
}
