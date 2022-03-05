# gjson
基于golang encoding/json开发的json工具库
## 工具一览
### ExtractValueFromJson
从json字符串中取出指定key的value值
```go
// 用法
distValue, err := gjson.ExtractValueFromJson(jsonStr, key)
```
其中`jsonStr`为json字符串，`key`为需要取出的值  
`key`说明: 嵌套的key请用`.`作为分隔，例如`{"root": {"sub": "distValue"}}`, 欲取出`sub`的值则输入`key`为`root.sub`  
`distValue`说明: 返回结果为`interface{}`, 需要自行转型, 若取出的值有多个值, 则类型为`[]interface{}`, 请注意做处理
```go
switch distValue.(type) {
case []interface{}:
    // 有多个返回值
default: 
    // 返回单个值
}
```