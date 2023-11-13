package utils

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func Test_main(t *testing.T) {
	// 示例数据
	//xmlData := `<root><name>John Doe</name><age>25</age></root>`
	//jsonData := `{"name": "John Doe", "age": 25}`
	//arrayData := `users[]=John&users[]=Jane&users[]=Doe&products[]=Apple&products[]=Banana`
	postData := "users=John&users1=Jane&users2=Doe&products=123"
	//fmt.Println(ReplaceXMLData(xmlData, "fuzz0"))
	//fmt.Println(ReplaceJSONData(jsonData, "fuzz0"))
	//fmt.Println(ReplaceArrayData(arrayData, "fuzz0"))
	//fmt.Println(ReplaceAfterAndData(arrayData, "fuzz0"))
	fmt.Println(ReplaceTagTypeMap["="].ReplaceTag(postData))
	strArr := []string{",", "'", "\"", "(", ")", "."}
	num := 10

	randomStrings := RandomStr(strArr, num)
	spew.Dump(randomStrings)

}
func ParseJson(t *testing.T) {
	byt := []byte(`{
       "num":6.13,
       "strs":["a","b"],
       "obj":{"foo":{"bar":"zip","zap":6}}
   }`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(interface{})
	fmt.Println(num)

}
