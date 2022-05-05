// This sample program demonstrates how to decode a JSON string.
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// JSON contains a sample string to unmarshal.
var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

func main() {
	// Unmarshal the JSON string into our map variable.
	// map[string]interface{} 值为interface{}类型,类型可以使用任意类型的值. 带来的问题见下
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Println(c)
	fmt.Println("Name:", c["name"])
	fmt.Println("Title:", c["title"])
	fmt.Println("Contact")
	fmt.Println("H:", c["contact"].(map[string]interface{})["home"]) //问题,所以必须将值转换为合适的类型，才能处理这个值
	fmt.Println("C:", c["contact"].(map[string]interface{})["cell"])

}
