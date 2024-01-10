package hangmanweb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Save(table [][]string) {
	encoded, err := json.Marshal(table)
	if err == nil {
		ioutil.WriteFile("database.json", encoded, 0777)
	} else {
		fmt.Println(err)
	}
}

func Load() [][]string {
	var Table [][]string
	data, _ := ioutil.ReadFile("database.json")
	dele := json.Unmarshal(data, &Table)
	if dele != nil {
		fmt.Println(dele)
	}
	return Table
}
