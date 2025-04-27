package data

import "fmt"

var Countries [10]string
var Slice []int
var Map map[int]string // key, value it like objects in JS

func init() {
	Countries[0] = "USA"
	Countries[1] = "Canada"
	Countries[2] = "Mexico"
	Countries[3] = "UK"
	Countries[4] = "Germany"
	Countries[5] = "France"
	Countries[6] = "Italy"
	Countries[7] = "Spain"
	Countries[8] = "Australia"
	Countries[9] = "Japan"

	qty := len(Countries)
	fmt.Println("Countries initialized", qty)

}
