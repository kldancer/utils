package dataStructure

import "github.com/tidwall/gjson"

func JsonUse() {
	//var json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	//value := gjson.Get(json, "name.last")

	var testjons = `{"Status":{"Conditions":{"Last Transition Time":"2024-03-:17:31Z","Status":"Pending"},"Last Transition Time":"2024-03-19T07:17:34Z","Status":"Running"},"Last Transition Time":"2024-03-19T07:18:06Z","Status":"Completed"}`
	value := gjson.Get(testjons, "Status.conditions.#(type==\"Complete\")#|#(Status==\"Completed\")#")

	// 用于从 JSON 文档中快速获取值的工具。它使用点语法路径（如 “name.last” 或 “age”）在 JSON 中搜索指定的路径。当找到值时，它会立即返回
	println(value.String())
}
