package load

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Item struct {
	Name string
	Text string
}

type CheatSheet struct {
	Title string
	Items []Item
}

func PrepData() []CheatSheet {
	fmt.Println("Loading json data")
	file, _ := ioutil.ReadFile("data.json")
	var dat []CheatSheet

	if err := json.Unmarshal([]byte(file), &dat); err != nil {
		panic(err)
	}

	return dat
}

func main() {

	PrepData()

}
